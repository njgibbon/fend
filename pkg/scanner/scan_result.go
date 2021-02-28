package scanner

import (
	"fmt"
	"strings"
	"time"
)

//ScanResult to return results from Scan
type ScanResult struct {
	Time         time.Duration
	Total        int
	Passed       int
	Failed       int
	SkippedDirs  int
	SkippedFiles int
	Errors       int
	ErrorPaths   []string
	FailedPaths  []string
}

//Output prints the results
func (sr ScanResult) Output() string {
	failedPaths := strings.Join(sr.FailedPaths, ", ")
	errorPaths := strings.Join(sr.ErrorPaths, ", ")
	resultsPart := "Results\n-----\n"
	failedPart := ""
	errorPart := ""
	output := resultsPart
	if len(sr.FailedPaths) != 0 {
		failedPart = "Failed\n-----\n[ " + failedPaths + " ]\n-----\n"
	}
	if len(sr.ErrorPaths) != 0 {
		errorPart = "Errors\n-----\n[ " + errorPaths + " ]\n-----\n"
	}
	statsPart := "Stats\n-----\nTime: " + fmt.Sprint(sr.Time) +
		"\nTotal Files Scanned: " + fmt.Sprint(sr.Total) +
		"\nPassed: " + fmt.Sprint(sr.Passed) + " - (" + fmt.Sprintf("%.2f", sr.passedPercent()) + "%)" +
		"\nFailed: " + fmt.Sprint(sr.Failed) + " - (" + fmt.Sprintf("%.2f", sr.failedPercent()) + "%)" +
		"\nSkipped Dirs: " + fmt.Sprint(sr.SkippedFiles) +
		"\nSkipped Files: " + fmt.Sprint(sr.SkippedFiles) +
		"\nErrors: " + fmt.Sprint(sr.Errors) + " - (" + fmt.Sprintf("%.2f", sr.errorPercent()) + "%)"
	output = resultsPart + failedPart + errorPart + statsPart
	return output
}

func (sr ScanResult) passedPercent() float32 {
	if sr.Passed == 0 || sr.Total == 0 {
		return 0
	}
	if sr.Total == 0 {
		return 100
	}
	return (float32(sr.Passed) / float32(sr.Total)) * 100
}

func (sr ScanResult) failedPercent() float32 {
	if sr.Failed == 0 || sr.Total == 0 {
		return 0
	}
	return (float32(sr.Failed) / float32(sr.Total)) * 100
}

func (sr ScanResult) errorPercent() float32 {
	if sr.Errors == 0 || sr.Total == 0 {
		return 0
	}
	return (float32(sr.Errors) / float32(sr.Total)) * 100
}
