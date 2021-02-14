package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Skip struct {
		File      []string `yaml:"file"`
		Folder    []string `yaml:"folder"`
		FileAll   []string `yaml:"file_all"`
		FolderAll []string `yaml:"folder_all"`
		Extension []string `yaml:"extension"`
	} `yaml:"skip"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(".fend.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(cfg)
	doIt(cfg)
	subDirToSkip := ".git"
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && (info.Name() == subDirToSkip) {
			//Skip
			fmt.Println(path, "Skip")
			return filepath.SkipDir
		} else if info.Name() == "." || info.IsDir() || info.Size() == 0 {
			//Skip
			fmt.Println(path, "Skip")
		} else {
			result, err := checkLineEnding(path)
			if err != nil {
				return err
			}
			fmt.Println(path, result)
		}
		fmt.Println(info.Size())
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func doIt(cfg *Config) {
	fmt.Println(cfg.Skip.Extension)
	fmt.Println(cfg.Skip.File)
	fmt.Println(cfg.Skip.FileAll)
	fmt.Println(cfg.Skip.Extension[0])
}

func checkLineEnding(fname string) (bool, error) {
	posixNewLine := "\n"
	//windowsNewLine := "\r\n"
	//newLine := "\r\n"
	file, err := os.Open(fname)
	if err != nil {
		return false, err
	}
	defer file.Close()
	buf := make([]byte, 1)
	stat, err := os.Stat(fname)
	//fmt.Print(stat.Size())
	start := stat.Size() - 1
	_, err = file.ReadAt(buf, start)
	if err == nil {
		//fmt.Printf("%s\n", buf)
	}
	//fmt.Print(buf)
	//s := string([]byte{buf})
	myString := string(buf)
	// fmt.Print(myString)
	// fmt.Print(posixNewLine)
	// b := []byte(posixNewLine)
	// fmt.Print(b)
	if myString == posixNewLine {
		return true, nil
	}
	return false, nil
}
