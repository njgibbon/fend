package scanner

import (
	"os"
	"path/filepath"
	"time"
)

// Scan will scan a given directory for none newline File Endings and return some stats
// It will take into account the skip configurations passed in
func Scan(cfg *ScanConfig, checkDir string) (*ScanResult, error) {
	scanResult := new(ScanResult)
	scanResult.Total = 0
	scanResult.Passed = 0
	scanResult.Failed = 0
	scanResult.SkippedFiles = 0
	scanResult.SkippedDirs = 0
	scanResult.Errors = 0
	scanResult.FailedExtensionSet = make(map[string]bool)
	startTime := time.Now()

	err := filepath.Walk(checkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			scanResult.Total++
			scanResult.Errors++
			scanResult.ErrorPaths = append(scanResult.ErrorPaths, filepath.ToSlash(path))
			return err
		}
		objName := info.Name()
		fileExtension := filepath.Ext(objName)
		if fileExtension == "" {
			fileExtension = objName
		}
		normalisedPath := filepath.ToSlash(path)
		if info.IsDir() {
			pathInSkipDir := contains(cfg.Skip.Dir, normalisedPath)
			nameInSkipDirAll := contains(cfg.Skip.DirAll, objName)
			if objName == ".git" {
				return filepath.SkipDir
			} else if nameInSkipDirAll == true {
				scanResult.SkippedDirs++
				return filepath.SkipDir
			} else if pathInSkipDir == true {
				scanResult.SkippedDirs++
				return filepath.SkipDir
			} else {
				//Move on, can't process dir but nothing special to do
			}
		} else {
			pathInSkipFile := contains(cfg.Skip.File, normalisedPath)
			nameInSkipFileAll := contains(cfg.Skip.FileAll, objName)
			fileExtInSkipExt := contains(cfg.Skip.Extension, fileExtension)
			if objName == "." {
				//Skip but don't record
			} else if info.Size() == 0 {
				scanResult.Total++
				scanResult.Failed++
				scanResult.FailedPaths = append(scanResult.FailedPaths, normalisedPath)
				scanResult.FailedExtensionSet[fileExtension] = true
			} else if nameInSkipFileAll == true {
				scanResult.SkippedFiles++
			} else if pathInSkipFile == true {
				scanResult.SkippedFiles++
			} else if fileExtInSkipExt == true {
				scanResult.SkippedFiles++
			} else {
				result, err := checkLineEnding(path)
				if err != nil {
					scanResult.Total++
					scanResult.Errors++
					scanResult.ErrorPaths = append(scanResult.ErrorPaths, normalisedPath)
					return err
				}
				if result == true {
					scanResult.Total++
					scanResult.Passed++
				} else {
					scanResult.Total++
					scanResult.Failed++
					scanResult.FailedPaths = append(scanResult.FailedPaths, normalisedPath)
					scanResult.FailedExtensionSet[fileExtension] = true
				}
			}
		}
		return nil
	})
	scanResult.Time = time.Since(startTime)
	if err != nil {
		return scanResult, err
	}
	return scanResult, nil
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
