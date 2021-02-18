package scanner

import (
	"fmt"
	"os"
	"path/filepath"
)

// Scan will scan a given directory for none newline File Endings and return some stats
// It will take into account the skip configurations passed in
func Scan(skipFile []string, skipFileAll []string, skipDir []string, skipDirAll []string, skipExtension []string, checkDir string) (int, int, int, int, int, []string, error) {
	passed := 0
	failed := 0
	skippedFiles := 0
	skippedDirs := 0
	errors := 0
	errorPaths := []string{}

	err := filepath.Walk(checkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errors++
			//TODO: Add to errorPaths slice
			return err
		}
		objName := info.Name()
		fileExtension := filepath.Ext(objName)
		normalisedPath := filepath.ToSlash(path)
		pathInSkipDir := contains(skipDir, normalisedPath)
		pathInSkipFile := contains(skipFile, normalisedPath)
		nameInSkipDirAll := contains(skipDirAll, objName)
		nameInSkipFileAll := contains(skipFileAll, objName)
		fileExtInSkipExt := contains(skipExtension, fileExtension)
		if objName == ".git" {
			return filepath.SkipDir
		} else if objName == "." {
			//Skip but don't record
		} else if info.IsDir() && (nameInSkipDirAll == true) {
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
				//TODO: Add to errorPaths slice
				fmt.Println(normalisedPath)
			}
		}
		return nil
	})
	if err != nil {
		return passed, failed, skippedDirs, skippedFiles, errors, errorPaths, err
	}
	return passed, failed, skippedDirs, skippedFiles, errors, errorPaths, nil
}

// checklineEnding checks whether a given file ends with a newline
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
