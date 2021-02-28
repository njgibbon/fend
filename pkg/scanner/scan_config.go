package scanner

//ScanConfig to pass in scan confguration
type ScanConfig struct {
	Skip struct {
		File      []string
		Dir       []string
		FileAll   []string
		DirAll    []string
		Extension []string
	}
}
