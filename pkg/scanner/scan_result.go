package scanner

//ScanResult is this
type ScanResult struct {
	Total        int
	Passed       int
	Failed       int
	SkippedDirs  int
	SkippedFiles int
	Errors       int
	ErrorPaths   []string
	FailedPaths  []string
}
