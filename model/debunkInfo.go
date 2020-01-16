package model


type DebunkInfo struct {
	Colour   string  `json:"colour"`
	Content  string	 `json:"content"`
	SendTime string  `json:"send_time"`
}

type  Debunks struct {
	Debunkid  string  `gorm:"debunk_id",json:"secret_id"`
	Content   string  `gorm:"content",json:"content"`
	SendTime  string  `gorm:"post_time",json:"send_time"`
	SenderSid string  `gorm:"sender_sid",json:"-"`
	Colour    string  `gorm:"colour",json:"colour"`
}


