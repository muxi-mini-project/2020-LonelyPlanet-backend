package model


func CraeteDebunk(debunk Debunks) (secretid string,error error) {
	if err := Db.Self.Model(&Debunks{}).Create(&debunk).Error ; err != nil {
		error = nil
	}
	var secret Debunks
	Db.Self.Model(&Debunks{}).Where(Debunks{Content:debunk.Content}).Find(&secret)
	secretid = secret.Debunkid
	return
}

func DeleteDebunk(secretid string) (err error) {
	if err := Db.Self.Model(&Debunks{}).Where(Debunks{Debunkid:secretid}).Delete(Debunks{}) ; err != nil {
		return nil
	}
	return err
}

func HistoryDebunk(uid string) (history []Debunks,err error) {
	if err = Db.Self.Model(&Debunks{}).Where(Debunks{SenderSid:uid}).Find(&history).Error ; err != nil {
		err = nil
	}
	return
}




