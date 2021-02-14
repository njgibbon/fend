package main

import (
    "fmt"
    // "io/ioutil"
    "log"
	"os"
	"path/filepath"
)

func main() {
    // files, err := ioutil.ReadDir("./")
    // if err != nil {
    //     log.Fatal(err)
    // }
 
    // for _, f := range files {
    //         fmt.Println(f.Name())
    // }

    // file, err := os.Open(".dotfile")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer func() {
    //     if err = file.Close(); err != nil {
    //         log.Fatal(err)
    //     }
    // }()


  	//b, err := ioutil.ReadAll(file)
 	//fmt.Print(b)
  	err := filepath.Walk(".",func(path string, info os.FileInfo, err error) error {
  		if err != nil {
	  		return err
  		}
		fmt.Println(path, info.Size())
		if info.IsDir()	{
			//Skip
		} else {
			checkLineEnding(path)
		}
  		return nil
	})
if err != nil {
  log.Println(err)
}

  //readFile(".dotfile")
}

func checkLineEnding(fname string) {
    file, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer file.Close()
	fmt.Print("a")
    buf := make([]byte, 14)
    stat, err := os.Stat(fname)
	//fmt.Print(stat)
	fmt.Print(stat.Size())
    start := stat.Size() - 14
    _, err = file.ReadAt(buf, start)
    if err == nil {
        fmt.Printf("%s\n", buf)
    }
	fmt.Print("b")
	fmt.Print(buf)
	fmt.Printf("%s\n", buf)
	fmt.Printf("%s\n", buf)
	fmt.Printf("%s\n", buf)
	//s := string([]byte{buf})
	myString := string(buf)
	fmt.Print(myString)
	
}