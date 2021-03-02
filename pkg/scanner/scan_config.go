package scanner

//ScanConfig to pass scan confiuration into Scan
type ScanConfig struct {
	Skip struct {
		File      []string
		Dir       []string
		FileAll   []string
		DirAll    []string
		Extension []string
	}
}
