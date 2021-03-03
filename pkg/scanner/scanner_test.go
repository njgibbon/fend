package scanner

import (
	"fmt"
	"testing"
)

// Tests are coupled with the relative test/data.. directory
func TestScanData0(t *testing.T) {
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
		fmt.Println(scanResult.Output())
		if scanResult.Total != 8 {
			t.Errorf("Total: Expected 8.")
		}
		if scanResult.Passed != 5 {
			t.Errorf("Passed: Expected 5.")
		}
		if scanResult.Failed != 3 {
			t.Errorf("Failed: Expected 3.")
		}
		if scanResult.SkippedDirs != 2 {
			t.Errorf("Skipped Dirs: Expected 2.")
		}
		if scanResult.SkippedFiles != 3 {
			t.Errorf("Skipped Files: Expected 3.")
		}
		if scanResult.Errors != 0 {
			t.Errorf("Errors: Expected 0.")
		}
	}
}

func TestScanData1(t *testing.T) {
	scanConfig := new(ScanConfig)

	scanResult, err := Scan(scanConfig, "test/data-1")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Not expecting Error.")
	} else {
		fmt.Println(scanResult)
		fmt.Println(scanResult.Output())
		if scanResult.Total != 1 {
			t.Errorf("Total: Expected 1.")
		}
		if scanResult.Passed != 1 {
			t.Errorf("Passed: Expected 1.")
		}
		if scanResult.Failed != 0 {
			t.Errorf("Failed: Expected 0.")
		}
	}
}

func TestScanUnknownPath(t *testing.T) {
	scanConfig := new(ScanConfig)

	scanResult, err := Scan(scanConfig, "unknown-path")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(scanResult)
		t.Errorf("Was expecting error due to unknown path to scan.")
	}
}
