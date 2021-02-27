package scanner

import (
	"fmt"
	"testing"
)

// Tests are coupled with the relative test/data.. directory
func TestScan(t *testing.T) {
	scanConfig := new(ScanConfig)
	scanConfig.Skip.File = []string{"test/data-0/skip_file.txt"}
	scanConfig.Skip.Dir = []string{"test/data-0/skip-dir"}
	scanConfig.Skip.FileAll = []string{"skip_file_all.txt"}
	scanConfig.Skip.DirAll = []string{"skip-dir-all"}
	scanConfig.Skip.Extension = []string{".skip", ".ignore"}

	scanResult, err := Scan(scanConfig, "test/data-0")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Not expecting Error.")
	} else {
		fmt.Println(scanResult)
		if scanResult.Passed != 5 {
			t.Errorf("Passed: Expected 5.")
		}

		//TODO Rest
	}
}

func TestScanUnknownPath(t *testing.T) {
	scanConfig := new(ScanConfig)
	scanConfig.Skip.File = []string{"test/data-0/skip_file.txt"}
	scanConfig.Skip.Dir = []string{"test/data-0/skip-dir"}
	scanConfig.Skip.FileAll = []string{"skip_file_all.txt"}
	scanConfig.Skip.DirAll = []string{"skip-dir-all"}
	scanConfig.Skip.Extension = []string{".skip", ".ignore"}

	scanResult, err := Scan(scanConfig, "unknown-path")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(scanResult)
		t.Errorf("Was expecting error due to unknown path to scan.")
	}
}
