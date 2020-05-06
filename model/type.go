package model

type UserInfo struct {
	Sid      string `json:"-" gorm:"sid"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	College  string `json:"college" gorm:"college"`
	Gender   string `json:"gender" gorm:"gender"`
	Grade    string `json:"grade" gorm:"grade"`
	Portrait int    `json:"portrait" gorm:"portrait"`
	Count    int    `json:"count" gorm:"default:0"`
	//NightPortrait string `json:"night_portrait" gorm:"night_portrait"`
	//Requirements  int    `json:"-" gorm:"requirements"`
	//Debunks       int    `json:"-" gorm:"debunks"`
}

type LoginInfo struct {
	Sid string `json:"sid"`
	Pwd string `json:"pwd"`
}

type AllNightRemindInfo struct {
	Titles   string `json:"titles"`
	PostTime string `json:"post_time"`
	Type2    string `json:"type_2"`
}

type Application struct {
	ContactWay []string `json:"contact_way"` //联系方式[qq, tel]
	Content    string   `json:"content"`     //附加信息
}

type AcceptApplication struct {
	ContactWay []string `json:"contact_way"` //联系方式[qq, tel]
	Content    string   `json:"content"`     //附加信息
}

type Requirements struct {
	RequirementId int    `gorm:"requirement_id" json:"requirement_id"`
	SenderSid     string `gorm:"sender_sid" json:"-"`
	Title         string `gorm:"title" json:"title"`
	Content       string `gorm:"content" json:"content"`
	PostTime      string `gorm:"post_time" json:"post_time"`
	Date          int    `gorm:"date" json:"date"`
	TimeFrom      int    `gorm:"time_form" json:"time_from"`
	TimeEnd       int    `gorm:"time_end" json:"time_end"`
	//RequirePeopleNum int    `gorm:"require_people_num" json:"require_people_num"`
	Place int `gorm:"place" json:"place"`
	Tag   int `gorm:"tag" json:"tag"`
	Type  int `gorm:"type" json:"type"`
	//ContactWayType   string `gorm:"contact_way_type" json:"contact_way_type"`
	//ContactWay       string `gorm:"contact_way" json:"contact_way"`
	Status int `gorm:"default:1" json:"-"`
	IsDraft  int    `gorm:"-" json:"is_draft"`  //新增，是否是草稿
}

type NewRequirements struct {
	Title    string `gorm:"title" json:"title"`
	Content  string `gorm:"content" json:"content"`
	PostTime string `gorm:"post_time" json:"post_time"`
	Date     int    `gorm:"date" json:"date"`
	TimeFrom int    `gorm:"time_form" json:"time_from"`
	TimeEnd  int    `gorm:"time_end" json:"time_end"`
	Place    int    `gorm:"place" json:"place"`
	Tag      int    `gorm:"tag" json:"tag"`
	Type     int    `gorm:"type" json:"type"`
}

type Res struct {
	Msg string `json:"msg"`
}

type ApplicationView struct {
	Msg          string                `json:"msg"`
	Num          int                   `json:"num"`
	Applications []ViewApplicationInfo `json:"applications"`
}

type Square struct {
	Msg     string                `json:"msg"`
	Num     int                   `json:"num"`
	Content []requirementInSquare `json:"content"`
}

type ViewRequirement struct {
	Msg     string      `json:"msg"`
	Content Requirement `json:"content"`
}

type ViewHistoryRequirement struct {
	Msg     string               `json:"msg"`
	Num     int                  `json:"num"`
	History []HistoryRequirement `json:"history"`
}

type RemindEx struct {
	Msg       string `json:"msg"`
	Existence bool   `json:"existence"`
}

type RemindBox struct {
	Msg     string         `json:"msg"`
	Num     int            `json:"num"`
	Content []ReminderInfo `json:"content"`
}

type ViewApplicationInfo struct {
	ApplicationId  int      `json:"application_id"`
	SenderNickname string   `json:"sender_nickname"`
	RequirementsId int      `json:"requirement_id"`
	College        string   `json:"college"`
	SendTime       string   `json:"send_time"`
	Title          string   `json:"title"`
	Gender         string   `json:"gender"`
	Grade          string   `json:"grade"`
	RedPoint       bool     `json:"red_point"`
	Portrait       int      `json:"portrait"`    //头像
	ContactWay     []string `json:"contact_way"` //联系方式[qq, tel]
	Content        string   `json:"content"`     //附加信息
}

type MyInformation struct {
	Msg      string `json:"msg"`
	Sid      string `json:"sid"`
	Nickname string `json:"nickname"`
	College  string `json:"college"`
	Gender   string `json:"gender"`
	Grade    string `json:"grade"`
	Portrait int    `json:"portrait"`
}

type ReportInformation struct {
	Reason string `json:"reason"`
	Addition string `json:"addition"`
}

type Draft struct {
	HasDraft int `json:"has_draft"`
	Content Requirements `json:"content"`
}

type FeedBackContent struct {
	Content string `json:"content"`
}