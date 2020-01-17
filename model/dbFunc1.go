package model

import (
	"strconv"
)

func CraeteDebunk(debunk Debunk) (secretid string,err error) {
	if err := Db.Self.Model(&Debunk{}).Create(&debunk).Error ; err != nil {
		err = nil
	}
	var secret Debunk
	Db.Self.Model(&Debunk{}).Where(Debunk{Content:debunk.Content}).Find(&secret)
	secretid = strconv.Itoa(secret.Debunkid)
	return
}

func DeleteDebunk(secretid int) (err error) {
	if err := Db.Self.Model(&Debunk{}).Where(Debunk{Debunkid:secretid}).Delete(Debunk{}) ; err != nil {
		return nil
	}
	return err
}

func HistoryDebunk(uid string) (history []Debunk,err error) {
	if err = Db.Self.Model(&Debunk{}).Where(Debunk{SenderSid:uid}).Find(&history).Error ; err != nil {
		err = nil
	}
	return
}

func SquareDebunk(page ,limit int)(secret []Debunk,err error){
	if err = Db.Self.Model(&Debunk{}).Limit(limit).Offset((page-1)*limit).Find(&secret).Error ; err != nil{
		err = nil
	}
	return
}

func CreateComment(comment Night_comment) (err error){
	if err = Db.Self.Model(&Night_comment{}).Create(&comment).Error ; err != nil {
		err = nil
	}
	return
}

func HistoryComment(secretid int) (history []Night_comment,err error){
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{SecretId:secretid}).Find(&history).Error ; err != nil {
		err = nil
		return
	}
	return
}

func DeleteComment(commentid int) (err error) {
	if err = Db.Self.Model(&Night_comment{}).Where(Night_comment{CommentId:commentid}).Delete(Night_comment{}).Error ; err != nil{
		err = nil
	}
	return
}

func GetSecretid(uid string)(secretid []int,err error){
	if err = Db.Self.Model(&Debunk{}).Where(Debunk{SenderSid:uid}).Pluck("debunk_id",&secretid).Error ; err != nil {
		err = nil
		return
	}
	return
}

func GetCommentData(secretid []int)(commentdata []Commentdata,err error){
	var Commentdata []Commentdata
	for _ ,data := range secretid {
		Commentdata = nil
		if err := Db.Self.Model(&Night_comment{}).Select("comment_time,comment,secret_id").Where(Night_comment{SecretId: data}).Scan(&Commentdata).Error ; err != nil {
			err = nil
		}
		for _ ,data1 := range Commentdata {
			commentdata = append(commentdata,data1)
		}
	}
	return
}




