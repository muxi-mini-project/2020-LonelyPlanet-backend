package model

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

/*
type reminders struct {
	RemindId     int    `gorm:"remind_id"`
	RemindInfoId int    `gorm:"remind_info_id"`
	ReceiverSid  string `gorm:"receiver_sid"`
	Type         int    `gorm:"type"`
	ReaderStatus int    `gorm:"reader_status"`
	Title        string `gorm:"title"`
	ReceiveTime  string `gorm:"receive_time"`
}
*/

type application struct {
	ApplicationId       int    `gorm:"application_id" json:"application_id"`
	ReceiverSid         string `gorm:"receiver_sid"　json:"-"`
	SenderSid           string `gorm:"sender_sid" json:"-"`
	RequirementId       int    `gorm:"column:requirements_id" json:"requirements_id"`
	Confirm             int    `gorm:"default:1" json:"-"`
	SenderReadStatus    int    `gorm:"default:1" json:"-"`
	ReceiverReadStatus  int    `gorm:"default:1" json:"-"`
	SendTime            string `gorm:"send_time" json:"send_time"`
	ConfirmTime         string `gorm:"confirm_time" json:"confirm_time"`
	Title               string `gorm:"title" json:"title"`
	SenderContactWay1   string `gorm:"sender_contact_way1" `  //申请人qq
	SenderContactWay2   string `gorm:"sender_contact_way2" `  //申请人tel
	ReceiverContactWay1 string `gorm:"receiver_contact_way1"` //发布需求者qq
	ReceiverContactWay2 string `gorm:"receiver_contact_way2"` //发布需求者tel
	Addition1           string `gorm:"addition1"`             //申请者附加信息
	Addition2           string `gorm:"addition2"`             //需求者附加信息
}

type latestAction struct {
	Sid        string `gorm:"sid"`
	LatestTime string `gorm:"latest_time"`
	RandNum    int    `gorm:"rand_num"`
}

func CreatUser(tmpUser UserInfo) error {
	var num int
	if Db.Self.Model(&UserInfo{}).Where(UserInfo{Sid: tmpUser.Sid}).Count(&num); num == 0 {
		if err := Db.Self.Model(&UserInfo{}).Create(&tmpUser).Error; err != nil {
			return err
		}
		var tmpAction latestAction
		tmpAction.Sid = tmpUser.Sid
		tmpAction.RandNum = getRandomNum()
		tmpAction.LatestTime = NowTimeStampStr()
		if err := Db.Self.Model(&latestAction{}).Create(&tmpAction).Error; err != nil {
			return err
		}
	}
	return nil
}

func FindUser(uid string) (UserInfo, error) {
	var tmpUser UserInfo
	if err := Db.Self.Model(&UserInfo{}).Where(UserInfo{Sid: uid}).Find(&tmpUser).Error; err != nil {
		return tmpUser, err
	}
	return tmpUser, nil
}

func VerifyInfo(uid string, verifyItem string, verifyInfo string) error {
	var tmpUser UserInfo
	if verifyItem == "Nickname" {
		tmpUser.NickName = verifyInfo
	}
	if err := Db.Self.Model(&UserInfo{}).Where(UserInfo{Sid: uid}).Update(&tmpUser).Error; err != nil {
		return err
	}
	return nil
}

/*
func RequirementFind(type1 int, date string, time_from int, time_end int, tag []int, place []int, limit int, offset int) []requirements {

	db := Db.Self

	db = db.Model(&requirements{}).Where(requirements{Type1:type1})

	var tmpRequirements []requirements

	if len(place) != 0 {
		for _,v := range place {
			db = db.Model(&requirements{}).Where(requirements{Place:v})
		}
	}

	if len(tag) != 0 {
		for _,v := range tag {
			db = db.Model(&requirements{}).Where(requirements{Tag:v})
		}
	}

	if date != "10000000" {
		db = db.Where("time_date & 11000000 > 1000000 ")
	}

	//db.Find(&tmpRequirements)

	//var result []requirements
	if time_from != 0 && time_end != 0 {
		/*for _,v := range tmpRequirements {
			if math.Abs(float64(v.Time_from - time_from)) <= 2 {
				if math.Abs(float64(v.Time_end-time_end)) <= 2 {
					result = append(result,v)
				}
			}
		}*/
/*
		if time_from == 1 {
			db = db.Where("time_from = 24").Or("time_from = 2")
		}
		if time_from == 24 {
			db = db.Where("time_from = 23").Or("time_from = 1")
		}
		if time_from != 1 && time_from != 24 {
			db = db.Where(" ABS(time_from - ?) <= 1",time_from)
		}
		if time_end == 1 {
			db = db.Where("time_end = 24").Or("time_end = 2")
		}
		if time_end == 24 {
			db = db.Where("time_end = 1").Or("time_end = 23")
		}
		if time_end != 1 && time_end != 24 {
			db = db.Where("ABS(time_end - ?) <= 1",time_end)
		}

	}

	db.Limit(limit).Offset(offset).Find(&tmpRequire
			db = db.Where("time_from = 23").Or("time_from = 1")
		}
		if time_from != 1 && time_from != 24 {
			db = db.Where(" ABS(time_from - ?) <= 1",time_from)
		}
		if time_end == 1 {
			db = db.Where("time_end = 24").Or("time_end = 2")
		}
		if time_end == 24 {
			db = db.Where("time_end = 1").Or("time_end = 23")
		}
		if time_end != 1 && time_end != 24 {
			db = db.Where("ABS(time_end - ?) <= 1",time_end)
		}

	}
	return tmpRequirements
}
*/

func RecordAction(uid string, num int, t string) error {
	var tmpAction latestAction
	tmpAction.LatestTime = t
	//	tmpAction.Sid = uid
	if num != -1 { //只有在需要更新的时候更新
		tmpAction.RandNum = num //getRandomNum()
	}
	if err := Db.Self.Model(&latestAction{}).Where("sid = ?", uid).Update(tmpAction).Error; err != nil {
		return err
	}
	return nil
}

func getRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20)
}

type requirementInSquare struct {
	RequirementId int    `json:"requirement_id"`
	Title         string `json:"title"`
	Date          string `json:"date"`
	Tag           string `json:"tag"`
	Place         string `json:"place"`
}

func RequirementFind(type1 int, sid string, date int, timeFrom int, timeEnd int, tag []int, place []int, limit int, offset int) ([]requirementInSquare, error) {

	db := Db.Self

	var tmpRequirements []Requirements
	var result []requirementInSquare

	var tmpRecord latestAction

	var wg sync.WaitGroup
	wg.Add(2)
	var err1 error
	go func() {
		if offset == 0 {
			tmpRecord.RandNum = getRandomNum()
			tmpRecord.LatestTime = NowTimeStampStr()
			err := RecordAction(sid, tmpRecord.RandNum, tmpRecord.LatestTime)
			err1 = err
		} else {
			err1 = Db.Self.Model(&latestAction{}).Where("sid = ?", sid).Find(&tmpRecord).Error
		}
		wg.Done()
	}()

	go func() {
		db = db.Model(&Requirements{}).Where(Requirements{Type: type1}).Where("sender_sid != ?", sid).Where("status = 1")

		if len(place) != 0 {
			var sql2 string
			for i, v := range place {
				if i == 0 {
					sql2 += "(place = " + strconv.Itoa(v)
					continue
				}
				sql2 += " or place = " + strconv.Itoa(v)
			}
			sql2 += ")"
			db = db.Model(&Requirements{}).Where(sql2)
		}

		if len(tag) != 0 {
			var sql1 string
			for i, v := range tag {
				if i == 0 {
					sql1 += "(tag = " + strconv.Itoa(v)
					continue
				}
				sql1 += " or tag = " + strconv.Itoa(v)
			}
			sql1 += ")"
			db = db.Model(&Requirements{}).Where(sql1)
		}

		/*
			if len(date) != 0 {
				db = db.Where("date & ? != 0 ", date)
			}
		*/
		if date != 0 {
			//tmp,_ := strconv.Atoi(date)
			db = db.Where("date & ? > 128 ", date)
		}

		//var result []requirements
		//改
		if timeFrom != 0 {
			db = db.Where(" time_from - ? >= 1", timeFrom)
		}
		if timeEnd != 0 {
			db = db.Where(" ? - time_end >= 1", timeEnd)
		}
		wg.Done()
	}()

	wg.Wait()

	if err1 != nil {
		return result, err1
	}

	db = db.Where("post_time < ?", tmpRecord.LatestTime) //确保分页准确

	db = db.Order("rand(" + strconv.Itoa(tmpRecord.RandNum) + ")")

	if err := db.Offset(offset).Limit(limit).Find(&tmpRequirements).Error; err != nil {
		log.Print("search err:")
		fmt.Println(err)
		return result, err
	}
	for _, v := range tmpRequirements {
		tmpResult := requirementInSquare{
			RequirementId: v.RequirementId,
			Title:         v.Title,
			Date:          dateImprove(Dec2BinStr(v.Date)),
			Tag:           tagImprove(v.Tag, v.Type),
			Place:         placeImprove(v.Place, v.Type),
		}
		result = append(result, tmpResult)
	}
	return result, nil
}

//真-->存在
func ConfirmRequirementExist(requirements Requirements) (error, bool) {
	var tmpRequirement []Requirements
	if err := Db.Self.Model(&Requirements{}).Where("sender_sid = ?", requirements.SenderSid).Find(&tmpRequirement).Error; err != nil {
		fmt.Println(err)
		return err, false
	}
	for _, v := range tmpRequirement {
		if v.Type == requirements.Type && v.Title == requirements.Title && v.Status != 2 && v.Tag == requirements.Tag && v.Place == requirements.Place {
			return nil, true
		}
	}
	return nil, false
}

func CreatRequirement(requirements Requirements) error {
	if err := Db.Self.Model(&Requirements{}).Create(&requirements).Error; err != nil {
		log.Print("CreatRequirement err")
		fmt.Print(err)
		return err
	}
	return nil
}

type Requirement struct {
	SenderNickName string `json:"sender_nick_name"`
	SenderPortrait int    `json:"sender_portrait"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	PostTime       string `json:"post_time"`
	Date           string `json:"date"`
	TimeFrom       int    `json:"time_from"`
	TimeEnd        int    `json:"time_end"`
	//RequirePeopleNum int    `json:"require_people_num"`
	Place string `json:"place"`
	Tag   string `json:"tag"`
	Type  string `json:"type"`
}

//假：还存在，即未被删除
func RequirementInfo(requirementId int) (Requirement, bool, error) {
	var tmpInfo Requirements
	var info Requirement
	if err := Db.Self.Model(&Requirements{}).Where(Requirements{RequirementId: requirementId}).Find(&tmpInfo).Error; err != nil {
		log.Print("RequirementInfo err")
		fmt.Println(err)
		return info, false, err
	}
	if tmpInfo.Status == 2 {
		return info, true, nil
	}
	userInfo, err := FindUser(tmpInfo.SenderSid)
	if err != nil {
		return info, false, err
	}
	info.Title = tmpInfo.Title
	info.Content = tmpInfo.Content
	info.Type = mainTypeImprove(tmpInfo.Type)
	info.Date = dateImprove(Dec2BinStr(tmpInfo.Date))
	info.Place = placeImprove(tmpInfo.Place, tmpInfo.Type)
	info.Tag = tagImprove(tmpInfo.Tag, tmpInfo.Type)
	info.SenderNickName = userInfo.NickName
	info.SenderPortrait = userInfo.Portrait
	info.TimeFrom = tmpInfo.TimeFrom
	info.TimeEnd = tmpInfo.TimeEnd
	info.PostTime = timestamp2json(tmpInfo.PostTime)
	//info.RequirePeopleNum = tmpInfo.RequirePeopleNum
	return info, false, nil
}

type HistoryRequirement struct {
	RequirementId int    `json:"requirement_id"`
	Title         string `json:"title"`
	PostTime      string `json:"post_time"`
	Tag           string `json:"tag"`
}

func HistoryRequirementFind(uid string, offset int, limit int) ([]HistoryRequirement, error) {
	var tmpResult []Requirements
	var result []HistoryRequirement
	if err := Db.Self.Model(&Requirements{}).Where(Requirements{SenderSid: uid}).Where("status = 1").Order("requirement_id desc").Offset(offset).Limit(limit).Scan(&tmpResult).Error; err != nil {
		log.Print("HistoryRequirementFind")
		fmt.Println(err)
		return result, err
	}
	for _, v := range tmpResult {
		tmpResult2 := HistoryRequirement{
			RequirementId: v.RequirementId,
			Title:         v.Title,
			PostTime:      timestamp2json(v.PostTime),
			Tag:           tagImprove(v.Tag, v.Type),
		}
		result = append(result, tmpResult2)
	}
	return result, nil
}

//int: 4: 申请的自己的需求 3: 已经申请过了 2: 申请的需求不存在 1: 成功 0: 未操作
func RequirementApply(uid string, requirementId int, contractWay1 string, contractWay2 string, addition string) (int, error) {
	tmpInfo, err := GetInfoFromRequirementId(requirementId)
	if err != nil {
		return 2, err
	}
	if tmpInfo.SenderSid == uid {
		return 4, nil
	}
	tmpApply := application{
		ReceiverSid:       tmpInfo.SenderSid,
		SenderSid:         uid,
		RequirementId:     requirementId,
		Confirm:           1,
		SendTime:          NowTimeStampStr(),
		Title:             tmpInfo.Title,
		SenderContactWay1: contractWay1,
		SenderContactWay2: contractWay2,
		Addition1:         addition, //附加信息
	}

	var num int
	if err := Db.Self.Model(&application{}).Where(application{SenderSid: uid, RequirementId: requirementId}).Count(&num).Error; err != nil {
		return 0, err
	}
	if num != 0 {
		return 3, nil
	}

	if err := Db.Self.Model(&application{}).Create(&tmpApply).Error; err != nil {
		log.Print("RequirementApply err")
		fmt.Println(err)
		return 0, err
	}

	//新增提醒

	return 1, nil
}

func RequirementDelete(requirementId int, sid string) (error, bool) {
	tmpInfo, err := GetInfoFromRequirementId(requirementId)
	if err != nil {
		return err, false
	}
	if tmpInfo.SenderSid != sid {
		return nil, false
	}
	var wg sync.WaitGroup
	wg.Add(2)
	var err1, err2 error
	go func() {
		if err := Db.Self.Model(&Requirements{}).Where(Requirements{RequirementId: requirementId}).Update(Requirements{Status: 2}).Error; err != nil {
			log.Print("RequirementDelete err")
			fmt.Println(err)
			err1 = err
		}
		wg.Done()
	}()

	go func() {
		if err := Db.Self.Model(&application{}).Where(application{RequirementId: requirementId}).Update(application{Confirm: 4}).Error; err != nil {
			log.Print("RequirementDelete err")
			fmt.Println(err)
			err2 = err
		}
		wg.Done()
	}()
	wg.Wait()
	if err1 != nil {
		return err1, true
	}
	if err2 != nil {
		return err2, true
	}
	return nil, true
}

func ConfirmRemindExist(uid string) (bool, error) {
	var num1, num int
	var err1, err2 error
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := Db.Self.Model(&application{}).Where("receiver_sid = ?", uid).Where("confirm = ?", 1).Where("receiver_read_status = ?", 1).Count(&num1).Error; err != nil {
			log.Print("ConfirmRemindExist:")
			fmt.Println(err)
			err1 = err
		}
		wg.Done()
	}()
	//有可能涉及删除还会提示的问题，应该是修复了，单需要进一步测试一下
	//同时，如果需要，可以返回一个提醒的数值
	go func() {
		if err := Db.Self.Model(&application{}).Where("sender_sid = ?", uid).Where("confirm != ?", 1).Where("confirm != ?", 4).Where("sender_read_status = ?", 1).Count(&num).Error; err != nil {
			log.Print("ConfirmRemindExist:")
			fmt.Println(err)
			err2 = err
		}
		wg.Done()
	}()
	wg.Wait()
	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}
	num = num1 + num
	if num == 0 {
		return false, nil
	}
	return true, nil
}

/*
func ViewAllUnsolvedApplication(uid string) ([]application, error) {
	db := Db.Self
	var tmpApplication []application
	if err := db.Where("receiver_sid = ?",uid).Where("confirm != ?",3).Find(&tmpApplication).Error; err != nil {
		log.Print("ViewAllUnsolvedApplication err")
		fmt.Println(err)
		return tmpApplication,err
	}
	return tmpApplication,nil
}
*/

func ViewAllApplication(uid string, offset int, limit int) ([]ViewApplicationInfo, error) {
	var tmpApplication []application
	var result []ViewApplicationInfo
	var tmpRecord latestAction
	if offset == 0 {
		tmpRecord.LatestTime = NowTimeStampStr()
		err := RecordAction(uid, -1, tmpRecord.LatestTime)
		if err != nil {
			return result, err
		}
	} else {
		if err := Db.Self.Model(&latestAction{}).Where("sid = ?", uid).Find(&tmpRecord).Error; err != nil {
			return result, err
		}
	}

	if err := Db.Self.Where("receiver_sid = ?", uid).Where("confirm != ?", 3).Where("confirm != ?", 4).Where("send_time < ?", tmpRecord.LatestTime).Order("application_id desc").Offset(offset).Limit(limit).Find(&tmpApplication).Error; err != nil {
		log.Print("ViewAllApplication err")
		fmt.Println(err)
		return result, err
	}

	for _, v := range tmpApplication {
		tmpUser, _ := FindUser(v.SenderSid)
		//tmpApplicationInfo := GetInfoFromRequirementId(v.RequirementsId)
		tmpResult := ViewApplicationInfo{
			ApplicationId:  v.ApplicationId,
			SenderNickname: tmpUser.NickName,
			RequirementsId: v.RequirementId,
			College:        changeCollegeName(tmpUser.College),
			SendTime:       timestamp2json(v.SendTime),
			Title:          v.Title,
			Gender:         tmpUser.Gender,
			Grade:          tmpUser.Grade,
			RedPoint:       redPoint(v.ReceiverReadStatus),
			Portrait:       tmpUser.Portrait,
			ContactWay:     []string{v.SenderContactWay1, v.SenderContactWay2},
			Content:        v.Addition1,
		}
		result = append(result, tmpResult)
	}
	//记得加数量
	return result, nil
}

type ViewApply struct {
	NikeName   string `gorm:"nike_name" json:"nike_name"`
	College    string `gorm:"college" json:"college"`
	Gender     string `gorm:"gender" json:"gender"`
	Grade      string `gorm:"grade" json:"grade"`
	ContactWay string `json:"contact_way"`
}

/*
func ViewApplication(applicationId int, uid string) (ViewApply, error) {//
	var tmpApplication application
	var result ViewApply
	if err := Db.Self.Model(&application{}).Where("application_id = ?",applicationId).Find(&tmpApplication).Error; err != nil {
		log.Print("ViewApplication err")
		fmt.Println(err)
		return result, err
	}
	//更新提醒列表，标为已读

	ReminderChangeStatus(applicationId, 1)//1="收到申请的人已读"
	return  , nil
}
*/

//int: 4 ->已删除 3 -> 非本人操作 2 ->已处理 1 -> 成功 0 -> 无操作
func SolveApplication(applicationId int, status int, sid string, addition AcceptApplication) (error, int) {
	var tmp application
	//排除了恶意application_id
	if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Find(&tmp).Error; err != nil {
		log.Print("SolveApplication err")
		fmt.Println(err)
		return err, 0
	}

	if tmp.Confirm == 4 {
		//return errors.New("是已经删除了的需求")
		return nil, 4
	}

	if tmp.Confirm == 2 {
		return nil, 2
	}
	tmp.ApplicationId = applicationId
	tmp.Confirm = status
	tmp.ConfirmTime = NowTimeStampStr()
	if status == 2 {
		tmp.ReceiverContactWay1 = addition.ContactWay[0]
		tmp.ReceiverContactWay2 = addition.ContactWay[1]
		tmp.Addition2 = addition.Content
	}

	if err := Db.Self.Model(&application{}).Where(application{ApplicationId: tmp.ApplicationId}).Update(tmp).Error; err != nil {
		log.Print("SolveApplication err")
		fmt.Println(err)
		return err, 0
	}
	err, flag := ReminderChangeStatus(applicationId, sid, 1)
	if err != nil {
		return err, 0
	}
	if !flag {
		return nil, 3
	}
	/*
		if err := Db.Self.Model(&application{}).Where(application{Application_id:application_id}).Delete(application{}).Error; err != nil {
			log.Print("SolveApplication delete err")
			fmt.Println(err)
			return err
		}
	*/
	/*
				if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Find(&tmp).Error; err != nil {
						log.Print("SolveApplication err")
					fmt.Println(err)
					return err
				}
				tmpInfo := GetInfoFromRequirementId(tmp.RequirementsId)
				newRemind := reminders{
					RemindInfoId: tmpInfo.RequirementsId,
					ReceiverSid:  tmpInfo.SenderSid,
					Type:         1,
					ReaderStatus: 0,
					Title:        tmpInfo.Title,
					ReceiveTime:  NowTime(),
				}
				if err := Db.Self.Model(&reminders{}).Create(&newRemind).Error; err != nil {
					log.Print("Creat new reminder fail")
					fmt.Println(err)
		func ReminderAdd(remindInfoId int, receiverId string, Tpye int, title string) {
			tmpReminder := reminders{
				RemindInfoId: remindInfoId,
				ReceiverSid:  receiverId,
				Type:         Tpye,
				ReaderStatus: 0,
				Title:        title,
				ReceiveTime:  NowTime(),
			}
			if err := Db.Self.Model(&reminders{}).Create(&tmpReminder).Error; err != nil {
				log.Print("ReminderAdd err")
				fmt.Println(err)
				return
			}
			return
		}

					return err
				}
	*/
	return nil, 1
}

type ReminderInfo struct {
	Status           int      `json:"status"`
	RequirementId    int      `json:"requirement_id"`
	Title            string   `json:"title"`
	ConfirmTime      string   `json:"confirm_time"`
	ContactWay       []string `json:"contact_way"` //需求者联系方式
	Content          string   `json:"content"`     //附加信息
	ReceiverNickName string   `json:"nick_name"`
	College          string   `json:"college"`
	Gender           string   `json:"gender"`
	Grade            string   `json:"grade"`
	RedPoint         bool     `json:"red_point"`
	ApplicationId    int      `json:"application_id"`
}

func ReminderBox(uid string, limit int, offset int) ([]ReminderInfo, error) {
	var result []ReminderInfo
	var tmp []application
	var tmpRecord latestAction

	if offset == 0 {
		tmpRecord.LatestTime = NowTimeStampStr()
		err := RecordAction(uid, -1, tmpRecord.LatestTime)
		if err != nil {
			return result, err
		}
	} else {
		if err := Db.Self.Model(&latestAction{}).Where("sid = ?", uid).Find(&tmpRecord).Error; err != nil {
			return result, err
		}
	}

	if err := Db.Self.Model(&application{}).Where("sender_sid = ?", uid).Where("confirm != ?", 1).Where("confirm != ?", 4).Where("confirm_time < ?", tmpRecord.LatestTime).Order("application_id desc").Offset(offset).Limit(limit).Find(&tmp).Error; err != nil {
		log.Print("ReminderBox err")
		fmt.Println(err)
		return result, nil
	}

	var tmpUserInfo UserInfo
	var tmpRequirement Requirements
	var err1, err2 error
	var wg sync.WaitGroup
	for _, v := range tmp {
		wg.Add(2)
		go func() {
			tmpUserInfo, err1 = FindUser(v.ReceiverSid)
			wg.Done()
		}()
		go func() {
			tmpRequirement, err2 = GetInfoFromRequirementId(v.RequirementId)
			wg.Done()
		}()
		wg.Wait()
		if err1 != nil {
			return result, err1
		}
		if err2 != nil {
			return result, err2
		}
		if v.Confirm == 2 {
			tmpInfo := ReminderInfo{
				Status:        v.Confirm, //通过是否来通知前端所显示的内容是否带有小眼睛图标
				RequirementId: v.RequirementId,
				Title:         v.Title,
				ConfirmTime:   timestamp2json(v.ConfirmTime),
				//ContactWayType:   tmpRequirement.ContactWayType,
				//ContactWay:       tmpRequirement.ContactWay,
				ContactWay:       []string{v.ReceiverContactWay1, v.ReceiverContactWay2},
				Content:          v.Addition2,
				ReceiverNickName: tmpUserInfo.NickName,
				College:          changeCollegeName(tmpUserInfo.College),
				Gender:           tmpUserInfo.Gender,
				Grade:            tmpUserInfo.Grade,
				RedPoint:         redPoint(v.SenderReadStatus), //控制小红点的显示
				ApplicationId:    v.ApplicationId,
			}
			result = append(result, tmpInfo)
		}
		if v.Confirm == 3 {
			tmpInfo := ReminderInfo{
				Status:           v.Confirm, //提示被拒绝啦！
				RequirementId:    v.RequirementId,
				Title:            v.Title,
				ReceiverNickName: tmpUserInfo.NickName,
				ConfirmTime:      timestamp2json(v.ConfirmTime),
				RedPoint:         redPoint(v.SenderReadStatus),
				ApplicationId:    v.ApplicationId,
			}
			result = append(result, tmpInfo)
		}
	}
	return result, nil
}

func redPoint(status int) bool {
	if status == 1 {
		return true
	}
	return false
}

//一些中间值转化
//后期检查
func GetInfoFromRequirementId(requirementID int) (Requirements, error) {
	var tmpInfo Requirements
	if err := Db.Self.Where(Requirements{RequirementId: requirementID}).Find(&tmpInfo).Error; err != nil {
		log.Print("GetInfoFromRequirementId err")
		fmt.Println(err)
		return tmpInfo, err
	}
	return tmpInfo, nil
}

//维护中间表
/*
func ReminderAdd(remindInfoId int, receiverId string, Tpye int, title string) {
	tmpReminder := reminders{
		RemindInfoId: remindInfoId,
		ReceiverSid:  receiverId,
		Type:         Tpye,
		ReaderStatus: 0,
		Title:        title,
		ReceiveTime:  NowTime(),
	}
	if err := Db.Self.Model(&reminders{}).Create(&tmpReminder).Error; err != nil {
		log.Print("ReminderAdd err")
		fmt.Println(err)
		return
	}
	return
}
*/

/*
func ReminderChangeStatus(remindInfoId int, uid string) {
	tmpReminder := reminders{RemindInfoId: remindInfoId, ReceiverSid: uid, ReaderStatus:1}
	if err := Db.Self.Model(&reminders{}).Where("remind_info_if = ?",remindInfoId).Where("receiver_sid = ?",uid).Update(&tmpReminder).Error; err != nil {
		log.Print("ReminderChangeStatus err")
		fmt.Println(err)
		return
	}
	return
}
*/

/*
func FindApplicationIdFromRequirementIdAndReceiverId(requirementId int, receiverId string) int {
	var applicationId int
	if err := Db.Self.Model(&application{}).Where(application{RequirementsId: requirementId, ReceiverSid: receiverId}).Pluck("application_id", &applicationId).Error; err != nil {
		log.Print("FindApplicationIdFromRequirementIdAndReceiverId")
		fmt.Println(err)
		return applicationId
	}
	return applicationId
}
*/

//更新阅读状态　type1 = 1 ->　收件人已读 type1 = 2 -> 发件人已读, true -> 本人 false ->　非本人/并不存在
func ReminderChangeStatus(applicationId int, sid string, type1 int) (error, bool) {
	if type1 == 1 {
		var num int
		if len(sid) != 0 {
			if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Where(application{ReceiverSid: sid}).Count(&num).Error; err != nil {
				return err, false
			}
			if num == 0 {
				return nil, false
			}
		}
		if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Update(application{ReceiverReadStatus: 2}).Error; err != nil {
			log.Print(" ReminderChangeStatus err ")
			fmt.Println(err)
			return err, true
		}
	}
	if type1 == 2 {
		var num int
		if len(sid) != 0 {
			if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Where(application{SenderSid: sid}).Count(&num).Error; err != nil {
				return err, false
			}
			if num == 0 {
				return nil, false
			}
		}
		if err := Db.Self.Model(&application{}).Where(application{ApplicationId: applicationId}).Update(application{SenderReadStatus: 2}).Error; err != nil {
			log.Print(" ReminderChangeStatus err ")
			fmt.Println(err)
			return err, false
		}
	}
	return nil, true
}
