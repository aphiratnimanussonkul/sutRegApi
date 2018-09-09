package model

// Datajsonn is dataRegSUT
type Datajsonn struct {
	Coursecode  string
	NameEnglish string
	NameThai    string
	Credit      string
	Datareg     []DReg
}

// DReg data each sec
type DReg struct {
	Sec     string
	DayTime []DayTime
	T       string
	Mid     Date
	Final   Date
	Note    string
}

//DayTime  day & time of each sec
type DayTime struct {
	Day  string
	Time string
}

// Date ex. 1-31
type Date struct {
	Date string
	Time string
}
