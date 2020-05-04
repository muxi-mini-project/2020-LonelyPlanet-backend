package model

import (
	//"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func CraeteDebunk(debunk Debunk) (secretid string, err error) {
	err = Db.Self.Model(&Debunk{}).Create(&debunk).Error
	if err != nil {
		return
	}
	var secret Debunk
	Db.Self.Model(&Debunk{}).Where(Debunk{Content: debunk.Content}).Find(&secret)
	secretid = strconv.Itoa(secret.Debunkid)
	return
}

func DeleteDebunk(secretid int) (err error) {
	if err := Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid: secretid}).Delete(Debunk{}).Error; err != nil {
		log.Println(err)
	}
	return err
}

func CheckDebunk(secretid int) bool {
	var data Debunk
	res := Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid: secretid}).Find(&data)
	if res.RecordNotFound() {
		return false
	}
	return true
}

func GetDebunk(secretid int) (data Debunk, err error) {
	if err := Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid: secretid}).Find(&data).Error; err != nil {
		log.Println(err)
	}
	return
}

func HistoryDebunk(uid string, page int, limit int) (history []Debunk, err error) {
	if err = Db.Self.Model(&Debunk{}).Where(Debunk{SenderSid: uid}).Limit(limit).Offset((page - 1) * limit).Find(&history).Error; err != nil {
		log.Println(err)
	}
	return
}

func RandNum(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}

func GetCommentHistory(history []Night_comment) (comment []Comment) {
	var comment1 Comment
	for _, data := range history {
		comment1 = Comment{
			Night_comment: data,
			Num:           RandNum(10),
		}
		comment = append(comment, comment1)
	}
	return
}

func SquareDebunk(sid string, page int, limit int) (secret Debunk, err error) {
	var i,n int
	var num int
	var tmpRecord latestAction
	var secretid []int

	if err = Db.Self.Model(&Debunk{}).Count(&i).Error; err != nil {
		log.Println(err)
		return
	}

	if page == 1 {
		tmpRecord.RandNum = getRandomNum()
		tmpRecord.LatestTime = NowTimeStampStr()
		err = RecordAction(sid, tmpRecord.RandNum, tmpRecord.LatestTime)
	}

	if i >= page * limit {
		if err = Db.Self.Model(&latestAction{}).Where("sid = ?", sid).Pluck("rand_num", &num).Error; err != nil {
			log.Println(err)
			return
		}
		if err = Db.Self.Model(&Debunk{}).Order("rand(" + strconv.Itoa(num) + ")").Limit(limit).Offset((page - 1) * limit).Find(&secret).Error; err != nil {
			log.Println(err)
			return
		}

	} else {
		if err = Db.Self.Model(&Debunk{}).Pluck("debunk_id", &secretid).Error; err != nil {
			log.Println(err)
			return
		}
		n = secretid[(RandNum(i))]
		if err = Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid: n}).Find(&secret).Error; err != nil {
			log.Println(err)
			return
		}
	}

	return
}

func CreateComment(comment Night_comment) (err error) {
	if err = Db.Self.Model(&Night_comment{}).Create(&comment).Error; err != nil {
		return err
	}
	return
}

func HistoryComment(secretid, page, limit int) (history []Night_comment, err error) {
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{SecretId: secretid}).Limit(limit).Offset((page - 1) * limit).Find(&history).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func HistoryComment1(secretid int) (history []Night_comment, err error) {
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{SecretId: secretid}).Find(&history).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func DeleteComment(commentid int) (err error) {
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId: commentid}).Delete(Night_comment{}).Error; err != nil {
		err = nil
	}
	return
}

func GetComment(commentid int) (comment Night_comment, err error)  {
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId: commentid}).Find(&comment).Error; err != nil {
		log.Println(err)
	}
	return
}

func CheckComment(commentid int) bool {
	var data Night_comment
	res := Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId: commentid}).Find(&data)
	if res.RecordNotFound() {
		return false
	}
	return true
}

func GetSecretid(uid string) (secretid []int, err error) {
	if err = Db.Self.Model(&Debunk{}).Where(Debunk{SenderSid: uid}).Pluck("debunk_id", &secretid).Error; err != nil {
		log.Println(err)
		return
	}
	return
}

func GetCommentData(uid string) (commentdata []Commentdata, err error) {
	var Commentdata1 []Night_comment
	var data2 Commentdata
	if err := Db.Self.Model(&Night_comment{}).Where(Night_comment{ReceiverSid: uid}).Order("comment_id desc").Find(&Commentdata1).Error; err != nil {
		log.Println(err)
	}
	for _, data1 := range Commentdata1 {
		data2.SecretId = data1.SecretId
		data2.CommentTime = data1.CommentTime
		data2.Comment = data1.Comment
		data2.CommentId = data1.CommentId
		data2.Status = data1.Status
		data2.Num = RandNum(10)
		commentdata = append(commentdata, data2)
	}
	return
}

func CheckRemindExist(uid string) (status int, err error) {
	var num int
	if err := Db.Self.Model(&Night_comment{}).Where(Night_comment{ReceiverSid: uid}).Where("status = ?", 0).Count(&num).Error; err != nil {
		log.Println(err)
	}
	if num == 0 {
		status = 1
	} else {
		status = 0
	}
	return
}

func CheckCommentIdExist(commentid int) (status int, err error) {
	var num int
	if err := Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId: commentid}).Count(&num).Error; err != nil {
		log.Println(err)
	}
	if num == 0 {
		status = 1
	} else {
		status = 0
	}
	return
}

func ChangeStatus(commentid int) (err error) {
	if err := Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId: commentid}).Update(Night_comment{Status: 1}).Error; err != nil {
		log.Println(err)
	}
	return
}
