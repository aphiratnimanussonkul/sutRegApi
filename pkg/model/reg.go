package model

// Datajsonn is dataRegSUT
type Datajsonn struct {
	Coursecode  string
	NameEnglish string
	NameThai    string
	Credit      string
	Datareg     []DReg
}
type DReg struct {
	Sec     string
	DayTime []DayTime
	T       string
	Mid     string
	Final   string
	Note    string
}

type DayTime struct {
	Day  string
	Time string
}
