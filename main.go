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
	defaultSkipDirAll  = []string{".git"}
	defaultSkipFileAll = []string{"."}
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
	fmt.Println("Fend - Check for Newline at File End\n-----\nConfig Loaded:", configLoaded,
		"\n-----\nScan - Files Failed\n-----")
	//Append default skip list to config list
	fendConfig.Skip.DirAll = append(fendConfig.Skip.DirAll, defaultSkipDirAll...)
	fendConfig.Skip.FileAll = append(fendConfig.Skip.FileAll, defaultSkipFileAll...)
	//fmt.Print(fendConfig)
	passed, failed, skippedDirs, skippedFiles, errors, err :=
		fend(fendConfig.Skip.File, fendConfig.Skip.FileAll, fendConfig.Skip.Dir, fendConfig.Skip.DirAll, fendConfig.Skip.Extension, ".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("-----\nResults\n-----")
	fmt.Println("Passed:", passed)
	fmt.Println("Failed:", failed)
	fmt.Println("Skipped Dirs:", skippedDirs)
	fmt.Println("Skipped Files:", skippedFiles)
	fmt.Println("Errors:", errors)
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

func fend(skipFile []string, skipFileAll []string, skipDir []string, skipDirAll []string, skipExtension []string, checkDir string) (int, int, int, int, int, error) {
	passed := 0
	failed := 0
	skippedFiles := 0
	skippedDirs := 0
	errors := 0

	err := filepath.Walk(checkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errors++
			return err
		}
		fileName := info.Name()
		fileExtension := filepath.Ext(fileName)
		//fmt.Println(fileExtension)
		normalisedPath := filepath.ToSlash(path)
		pathInSkipDir := contains(skipDir, normalisedPath)
		pathInSkipFile := contains(skipFile, normalisedPath)
		nameInSkipDirAll := contains(skipDirAll, fileName)
		nameInSkipFileAll := contains(skipFileAll, fileName)
		fileExtInSkipExt := contains(skipExtension, fileExtension)
		if info.IsDir() && (nameInSkipDirAll == true) {
			//fmt.Println(normalisedPath, "Skip - SkipDirAll")
			skippedDirs++
			return filepath.SkipDir
		} else if info.IsDir() && (pathInSkipDir == true) {
			//fmt.Println(normalisedPath, "Skip - SkipDir !!!!!!!!!!!!!!!!!!!!!!!!!")
			skippedDirs++
			return filepath.SkipDir
		} else if nameInSkipFileAll == true {
			//fmt.Println(normalisedPath, "Skip - SkipFileAll")
			skippedFiles++
		} else if info.IsDir() {
			//fmt.Println(normalisedPath, "Dir NA")
			//Move on, can't process folder but nothing special to do
		} else if info.Size() == 0 {
			fmt.Println(normalisedPath)
			failed++
		} else if pathInSkipFile == true {
			//fmt.Println(normalisedPath, "Skip - SkipFile")
			skippedFiles++
		} else if fileExtInSkipExt == true {
			//fmt.Println(normalisedPath, "Skip - Extension")
			skippedFiles++
		} else {
			result, err := checkLineEnding(path)
			if err != nil {
				errors++
				return err
			}
			if result == true {
				passed++
			} else {
				failed++
				fmt.Println(normalisedPath)
			}
		}
		return nil
	})
	if err != nil {
		return passed, failed, skippedDirs, skippedFiles, errors, err
	}
	return passed, failed, skippedDirs, skippedFiles, errors, nil
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

// contains is a helper function
// Does a slice contain this string?
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
