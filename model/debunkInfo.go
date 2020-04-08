package model

type DebunkInfo struct {
	Colour   string `json:"colour"`
	Content  string `json:"content"`
	SendTime string `json:"send_time"`
}

type Debunk struct {
	Debunkid  int    `gorm:"column:debunk_id",gorm:"AUTO_INCREMENT",json:"secret_id",json:"-"`
	Content   string `gorm:"column:content",json:"content"`
	SendTime  string `gorm:"column:post_time",json:"send_time"`
	SenderSid string `gorm:"column:sender_sid",json:"-"`
	Colour    string `gorm:"column:colour",json:"colour"`
}

type Night_comment struct {
	CommentTime  string `gorm:"column:comment_time",json:"-"`
	Comment      string `gorm:"column:content",json:"comment"`
	SecretId     int    `gorm:"column:debunk_id",json:"secret_id"`
	CommentId    int    `gorm:"column:comment_id",gorm:"AUTO_INCREMENT",json:"-"`
	ReceiverSid  string `gorm:"column:receiver_sid",gorm:"default:0",json:"receiver_sid"`
	Status		 int	`gorm:"column:status",json:"-"`
}

type Commentdata struct {
	CommentTime string `gorm:"column:comment_time",json:"comment_time"`
	Comment     string `gorm:"column:comment",json:"comment"`
	SecretId    int    `gorm:"column:secret_id",json:"secret_id"`
	CommentId   int    `gorm:"column:comment_id",gorm:"AUTO_INCREMENT",json:"-"`
	Status		int	   `gorm:"column:status",json:"-"`
	Num         int
}

type Comment struct {
	Night_comment
	Num int
}
