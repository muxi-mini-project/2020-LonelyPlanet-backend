package model


type DebunkInfo struct {
	Colour   string  `json:"colour"`
	Content  string	 `json:"content"`
	SendTime string  `json:"send_time"`
}

type  Debunk struct {
	Debunkid  int     `gorm:"debunk_id",json:"secret_id"`
	Content   string  `gorm:"content",json:"content"`
	SendTime  string  `gorm:"post_time",json:"send_time"`
	SenderSid string  `gorm:"sender_sid",json:"-"`
	Colour    string  `gorm:"colour",json:"colour"`
}

type Night_comment struct {
	CommentTime  string `gorm:"comment_time",json:"comment_time"`
	Comment      string `gorm:"content",json:"comment"`
	SecretId     int    `gorm:"debunk_id",json:"-"`
	CommentId    int    `gorm:"comment_id",json:"comment_id"`
}

type Commentdata struct {
	CommentTime  string  `gorm:"comment_time",json:"comment_time"`
	Comment      string  `gorm:"comment",json:"comment"`
	SecretId     int     `gorm:"secret_id",json:"secret_id"`
}


