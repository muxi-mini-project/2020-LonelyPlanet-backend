package model

import (
	"fmt"
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

func SquareDebunk() (secret Debunk, err error) {
	var i int
	var secretid []int
	if err = Db.Self.Model(&Debunk{}).Count(&i).Pluck("debunk_id", &secretid).Error; err != nil {
		log.Println(err)
		return
	}
	fmt.Println(i)
	n := secretid[(RandNum(i))]
	if err = Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid: n}).Find(&secret).Error; err != nil {
		log.Println(err)
		return
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

func GetCommentData(secretid []int) (commentdata []Commentdata, err error) {
	var Commentdata1 []Night_comment
	var data2 Commentdata
	for _, data := range secretid {
		Commentdata1 = nil
		if err := Db.Self.Model(&Night_comment{}).Where(Night_comment{SecretId: data}).Find(&Commentdata1).Error; err != nil {
			log.Println(err)
		}
		for _, data1 := range Commentdata1 {
			data2.SecretId = data1.SecretId
			data2.CommentTime = data1.CommentTime
			data2.Comment = data1.Comment
			data2.Num = RandNum(10)
			commentdata = append(commentdata, data2)
		}
	}
	return
}
