package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Constants for .git and .config
// Vars to use for version command

// Ouput: Path - Fail - Reason
// Summary: Stats
// Scanned: 67
// Skipped_Dirs: 22
// Skipped Files:
// Failed: 5
// Passed: 45

// Function for cross-platform path comparison.

// FendConfig is data for Fend Configuration annotated to be pulled from .fend.yaml
type FendConfig struct {
	Skip struct {
		File      []string `yaml:"file"`
		Folder    []string `yaml:"folder"`
		FileAll   []string `yaml:"file_all"`
		FolderAll []string `yaml:"folder_all"`
		Extension []string `yaml:"extension"`
	} `yaml:"skip"`
}

// newFendConfig returns a new decoded Config struct using .fend.yaml if it exists
func newFendConfig() (*FendConfig, error) {
	fendConfig := &FendConfig{}
	file, err := os.Open(".fend.yaml")
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

func main() {
	fmt.Println("Fend - Check for Newline at File End\n-----")
	fendConfig, err := newFendConfig()
	if err != nil {
		//Could not load .fend.yaml config file for some reason
		fmt.Println("Could not load .fend.yaml")
	}
	//Decision to always skip the .git dir
	fendConfig.Skip.FileAll = append(fendConfig.Skip.FileAll, ".git")
	fmt.Print(fendConfig)
	doIt(fendConfig)
	err = fend(fendConfig, ".")
	if err != nil {
		fmt.Println("Could not load .fend.yaml")
	}
}

func doIt(cfg *FendConfig) {
	fmt.Println(cfg.Skip.Extension)
	fmt.Println(cfg.Skip.File)
	fmt.Println(cfg.Skip.FileAll)
	//fmt.Println(cfg.Skip.Extension[0])
}

func fend(fendConfig *FendConfig, checkDir string) error {
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

func checkLineEnding(fname string) (bool, error) {
	posixNewLine := "\n"
	file, err := os.Open(fname)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 1)
	stat, err := os.Stat(fname)
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
