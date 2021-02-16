package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Constants for Fend
const (
	Version    = "dev"
	Source     = "https://github.com/njgibbon/fend"
	ConfigPath = ".fend.yaml"
)

var (
	defaultSkipDirAll = []string{".git"}
)

func main() {
	configLoaded := true
	flag.Parse()
	if flag.Arg(0) == "version" {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag.Arg(0) == "doc" || flag.Arg(0) == "help" {
		fmt.Println(Source)
		os.Exit(0)
	}
	fendConfig, err := newFendConfig(ConfigPath)
	if err != nil {
		configLoaded = false
	}
	fmt.Println("Fend - Check for Newline at File End\n-----")
	//Decision to always skip the .git dir
	fendConfig.Skip.FileAll = append(fendConfig.Skip.FileAll, ".git")
	fmt.Print(fendConfig)
	err = fend(fendConfig, ".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(configLoaded)
}

// FendConfig is data for Fend Configuration annotated to be pulled from .fend.yaml
type FendConfig struct {
	Skip struct {
		File      []string `yaml:"file"`
		Dir       []string `yaml:"dir"`
		FileAll   []string `yaml:"file_all"`
		DirAll    []string `yaml:"dir_all"`
		Extension []string `yaml:"extension"`
	} `yaml:"skip"`
}

// newFendConfig returns a new decoded FendConfig struct using Config File if exists
func newFendConfig(configPath string) (*FendConfig, error) {
	fendConfig := &FendConfig{}
	file, err := os.Open(configPath)
	if err != nil {
		return fendConfig, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&fendConfig); err != nil {
		return fendConfig, err
	}
	return fendConfig, nil
}

func fend(fendConfig *FendConfig, checkDir string) error {
	fmt.Println(fendConfig.Skip.Extension)
	fmt.Println(fendConfig.Skip.File)
	fmt.Println(fendConfig.Skip.FileAll)
	fmt.Println(defaultSkipDirAll)
	err := filepath.Walk(checkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//Main logic of whether to check
		if info.IsDir() && (info.Name() == ".git") {
			fmt.Println(path, "Skip .git")
			return filepath.SkipDir
		} else if info.Name() == "." {
			fmt.Println(path, ". NA")
		} else if info.IsDir() {
			fmt.Println(path, "Dir NA")
		} else if info.Size() == 0 {
			fmt.Println(path, "Size 0 - False")
		} else {
			result, err := checkLineEnding(path)
			if err != nil {
				return err
			}
			fmt.Println(path, result)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func checkLineEnding(fileName string) (bool, error) {
	posixNewLine := "\n"
	file, err := os.Open(fileName)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 1)
	stat, err := os.Stat(fileName)
	start := stat.Size() - 1
	_, err = file.ReadAt(buf, start)
	if err != nil {
		return false, err
	}
	lastCharacter := string(buf)
	if lastCharacter == posixNewLine {
		return true, nil
	}
	return false, nil
}