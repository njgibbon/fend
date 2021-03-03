package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/njgibbon/fend/pkg/scanner"
	"gopkg.in/yaml.v2"
)

// Constants for Fend
const (
	Version    = "1.0.0-rc"
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
	fmt.Println("Fend - Check for Newline at File End\n-----")
	cfg, err := newConfig(ConfigPath)
	if err != nil {
		configLoaded = false
	}
	scanConfig := newScanConfig(cfg)
	fmt.Println("Settings\n-----\nConfig Loaded:", configLoaded, "\n-----\nScan\n-----")
	scanResult, err := scanner.Scan(scanConfig, ".")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(scanResult.Output())
	if scanResult.Failed != 0 {
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
