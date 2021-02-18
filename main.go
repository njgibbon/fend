package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/njgibbon/fend/scanner"
	"gopkg.in/yaml.v2"
)

// Constants for Fend
const (
	Version    = "dev"
	Source     = "https://github.com/njgibbon/fend"
	ConfigPath = ".fend.yaml"
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
	cfg, err := newConfig(ConfigPath)
	if err != nil {
		configLoaded = false
	}
	fmt.Println("Fend - Check for Newline at File End\n-----\nConfig Loaded:", configLoaded,
		"\n-----\nScan\n-----")
	passed, failed, skippedDirs, skippedFiles, errors, errorPaths, err :=
		scanner.Scan(cfg.Skip.File, cfg.Skip.FileAll, cfg.Skip.Dir, cfg.Skip.DirAll, cfg.Skip.Extension, ".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Errors:", errorPaths)
	fmt.Println("-----\nResults\n-----")
	fmt.Println("Passed:", passed)
	fmt.Println("Failed:", failed)
	fmt.Println("Skipped Dirs:", skippedDirs)
	fmt.Println("Skipped Files:", skippedFiles)
	fmt.Println("Errors:", errors)
	if failed != 0 {
		os.Exit(1)
	}
}

// Config is data for Fend Configuration annotated to be pulled from .fend.yaml
type Config struct {
	Skip struct {
		File      []string `yaml:"file"`
		Dir       []string `yaml:"dir"`
		FileAll   []string `yaml:"file_all"`
		DirAll    []string `yaml:"dir_all"`
		Extension []string `yaml:"extension"`
	} `yaml:"skip"`
}

// newConfig returns a new decoded Config struct using Config File if exists
func newConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
