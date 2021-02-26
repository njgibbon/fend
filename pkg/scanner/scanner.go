package scanner

import (
	"os"
	"path/filepath"
)

// Scan will scan a given directory for none newline File Endings and return some stats
// It will take into account the skip configurations passed in
func Scan(skipFile []string, skipFileAll []string, skipDir []string, skipDirAll []string, skipExtension []string, checkDir string) (int, int, int, int, int, []string, []string, error) {
	passed := 0
	failed := 0
	skippedFiles := 0
	skippedDirs := 0
	errors := 0
	errorPaths := []string{}
	failedPaths := []string{}

	err := filepath.Walk(checkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			errors++
			errorPaths = append(errorPaths, filepath.ToSlash(path))
			return err
		}
		objName := info.Name()
		fileExtension := filepath.Ext(objName)
		normalisedPath := filepath.ToSlash(path)
		if info.IsDir() {
			pathInSkipDir := contains(skipDir, normalisedPath)
			nameInSkipDirAll := contains(skipDirAll, objName)
			if objName == ".git" {
				return filepath.SkipDir
			} else if nameInSkipDirAll == true {
				skippedDirs++
				return filepath.SkipDir
			} else if pathInSkipDir == true {
				skippedDirs++
				return filepath.SkipDir
			} else {
				//Move on, can't process dir but nothing special to do
			}
		} else {
			pathInSkipFile := contains(skipFile, normalisedPath)
			nameInSkipFileAll := contains(skipFileAll, objName)
			fileExtInSkipExt := contains(skipExtension, fileExtension)
			if objName == "." {
				//Skip but don't record
			} else if info.Size() == 0 {
				failed++
			} else if nameInSkipFileAll == true {
				skippedFiles++
			} else if pathInSkipFile == true {
				skippedFiles++
			} else if fileExtInSkipExt == true {
				skippedFiles++
			} else {
				result, err := checkLineEnding(path)
				if err != nil {
					errors++
					errorPaths = append(errorPaths, normalisedPath)
					return err
				}
				if result == true {
					passed++
				} else {
					failed++
					failedPaths = append(failedPaths, normalisedPath)
				}
			}
		}
		return nil
	})
	if err != nil {
		return passed, failed, skippedDirs, skippedFiles, errors, errorPaths, failedPaths, err
	}
	return passed, failed, skippedDirs, skippedFiles, errors, errorPaths, failedPaths, nil
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
