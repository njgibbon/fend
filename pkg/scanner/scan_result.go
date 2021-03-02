package scanner

import (
	"fmt"
	"time"
)

//ScanResult to return results from Scan
type ScanResult struct {
	Time               time.Duration
	Total              int
	Passed             int
	Failed             int
	SkippedDirs        int
	SkippedFiles       int
	Errors             int
	ErrorPaths         []string
	FailedPaths        []string
	FailedExtensionSet map[string]bool
}

//Output formats all of the results
func (sr ScanResult) Output() string {
	failedExtensionSet := []string{}
	for key := range sr.FailedExtensionSet {
		failedExtensionSet = append(failedExtensionSet, key)
	}
	resultsPart := "Results\n-----\n"
	failedPart := ""
	errorPart := ""
	output := resultsPart
	if len(sr.FailedPaths) != 0 {
		failedPart = "Failed\n-----\n" + fmt.Sprint(sr.FailedPaths) + "\n-----\n" +
			fmt.Sprint(failedExtensionSet) + "\n-----\n"
	}
	if len(sr.ErrorPaths) != 0 {
		errorPart = "Errors\n-----\n" + fmt.Sprint(sr.ErrorPaths) + "\n-----\n"
	}
	statsPart := "Stats\n-----\nTime: " + fmt.Sprint(sr.Time) +
		"\nTotal Files Scanned: " + fmt.Sprint(sr.Total) +
		"\nPassed: " + fmt.Sprint(sr.Passed) + " - (" + fmt.Sprintf("%.2f", sr.passedPercent()) + "%)" +
		"\nFailed: " + fmt.Sprint(sr.Failed) + " - (" + fmt.Sprintf("%.2f", sr.failedPercent()) + "%)" +
		"\nSkipped Dirs: " + fmt.Sprint(sr.SkippedDirs) +
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
