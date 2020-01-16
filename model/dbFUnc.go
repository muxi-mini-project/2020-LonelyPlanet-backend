package model


func CreatUser(tmpUser UserInfo) error {
	if err := Db.Self.Model(&UserInfo{}).Create(&tmpUser).Error; err != nil {
		return err
	}
	return nil
}

func FindUser(uid string) (UserInfo, error) {
	var tmpUser UserInfo
	if err := Db.Self.Model(&UserInfo{}).Where(UserInfo{Sid:uid}).Find(&tmpUser).Error; err != nil {
		return tmpUser,err
	}
	return tmpUser,nil
}

func VerifyInfo(uid string, verifyItem string, verifyInfo string) error {
	var tmpUser UserInfo
	if verifyItem == "Nickname" {
		tmpUser.NickName = verifyInfo
	}
	if err := Db.Self.Model(&UserInfo{}).Where(UserInfo{Sid:uid}).Update(tmpUser).Error; err != nil {
		return err
	}
	return nil
}
