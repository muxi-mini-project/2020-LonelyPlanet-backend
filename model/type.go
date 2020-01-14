package model


type UserInfo struct {
	Sid           string `json:"-" gorm:"sid"`
	NickName      string `json:"nick_name" gorm:"nick_name"`
	College       string `json:"college" gorm:"college"`
	Gender        string `json:"gender" gorm:"gender"`
	Grade         string `json:"grade" gorm:"grade"`
	NightPortrait string `json:"night_portrait" gorm:"night_portrait"`
	Requirements  int    `json:"-" gorm:"requirements"`
	Debunks       int    `json:"-" gorm:"debunks"`
}

type LoginInfo struct {
	Sid string `json:"sid"`
	Pwd string `json:"pwd"`
}

