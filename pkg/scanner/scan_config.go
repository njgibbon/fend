package scanner

//ScanConfig is this
type ScanConfig struct {
	Skip struct {
		File      []string
		Dir       []string
		FileAll   []string
		DirAll    []string
		Extension []string
	}
}
