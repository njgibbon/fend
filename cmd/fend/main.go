package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/njgibbon/fend/pkg/scanner"
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
	var target string
	flag.StringVar(&target, "target", "", "Target scan against a given file extension or other suffix.")
	flag.Parse()
	if flag.Arg(0) == "version" {
		fmt.Println(Version)
		os.Exit(0)
	}
	if flag.Arg(0) == "doc" || flag.Arg(0) == "help" {
		fmt.Println(Source)
		os.Exit(0)
	}
	fmt.Println("Fend - Check for Newline at File End\n-----\nSettings\n-----")
	if target != "" {
		fmt.Println("Mode: Target=", target, "\nConfig Loaded: n/a\n-----\nScan\n-----")
		os.Exit(0)
	} else {
		cfg, err := newConfig(ConfigPath)
		if err != nil {
			configLoaded = false
		}
		scanConfig := newScanConfig(cfg)
		fmt.Println(scanConfig)
		fmt.Println("Mode: Normal\nConfig Loaded:", configLoaded, "\n-----\nScan\n-----")
		start := time.Now()
		scanResult, err := scanner.Scan(scanConfig, ".")
		duration := time.Since(start)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Failed\n-----\n", scanResult.FailedPaths, "\n-----")
		fmt.Println("Errors\n-----\n", scanResult.ErrorPaths, "\n-----")
		fmt.Println("Results\n-----")
		fmt.Println("Time:", duration)
		fmt.Println("Total Files Scanned:", scanResult.Total)
		fmt.Println("Passed:", scanResult.Passed)
		fmt.Println("Failed:", scanResult.Failed)
		fmt.Println("Skipped Dirs:", scanResult.SkippedDirs)
		fmt.Println("Skipped Files:", scanResult.SkippedFiles)
		fmt.Println("Errors:", scanResult.Errors)
		if scanResult.Failed != 0 {
			os.Exit(1)
		}
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

// newScanConfig returns the scanner Type needed for Scan using the YAML Loaded Struct
func newScanConfig(config *Config) *scanner.ScanConfig {
	scanConfig := new(scanner.ScanConfig)
	scanConfig.Skip.File = config.Skip.File
	scanConfig.Skip.Dir = config.Skip.Dir
	scanConfig.Skip.FileAll = config.Skip.FileAll
	scanConfig.Skip.DirAll = config.Skip.DirAll
	scanConfig.Skip.Extension = config.Skip.Extension
	return scanConfig
}
