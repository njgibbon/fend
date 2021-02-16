package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/njgibbon/fend/fend"
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
		fend.FendScan(fendConfig.Skip.File, fendConfig.Skip.FileAll, fendConfig.Skip.Dir, fendConfig.Skip.DirAll, fendConfig.Skip.Extension, ".")
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
