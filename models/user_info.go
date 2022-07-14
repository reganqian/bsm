package models

import (
	. "common/static"
)

type UserInfo struct {
	UserId uint32 `gorm:"primary_key" json:"userId"`
	UserName string `json:"userName"`
	UserPhone string `json:"userPhone"`
	UserPwd string `json:"userPwd"`
	UserSalt string `json:"userSalt"`
	UserStatus string `json:"userStatus"`
	CreateTime int64 `json:"createTime"`
	LastLoginTime int64 `json:"lastLoginTime"`
	LastLoginIp string `json:"lastLoginIp"`
	LastLoginServer string `json:"lastLoginServer"`
	AuthStatus string `json:"authStatus"`
	UserCard string `json:"userCard"`
	Sex string `json:"sex"`
	RealName string `json:"realName"`
	SelfDesc string `json:"selfDesc"`
	Balance float64 `json:"balance"`
	OnlineStatus  string `json:"onlineStatus"`
	UserType string `json:"userType"`
	UserHead string `json:"userHead"`
	ContactQQ string `json:"contactQQ"`
	// contact_qq
}


//根据手机号从查询账号信息
func QueryByPhone(userPhone string) (data UserInfo, err error) {
	err = db.Table("user_info_tb").Where("user_phone = ?", userPhone).First(&data).Error
	return data, err
}


func SaveUserInfo(userPhone, userPwd, userSalt string) (userId uint32, err error) {
	data := UserInfo{}
	data.UserPhone = userPhone
	data.UserPwd = userPwd
	data.UserSalt = userSalt
	data.UserType = USERTYPE_NORMAL
	err = db.Table("user_info_tb").Create(&data).Error
	userId = data.UserId
	return userId, err
}

//根据id查询获取用户信息
func GetInfoById(userId uint32) (data UserInfo, err error) {
	err = db.Table("user_info_tb").Where("user_id = ?", userId).First(&data).Error
	return data, err
}

//统计手机号已有数量， 校验是否已存在
func GetCountByPhone(phoneNum string) (num int, err error) {
	err = db.Table("user_info_tb").Where("user_phone = ?", phoneNum).Count(&num).Error
	return num, err
} 

func UpdatePwd(userId uint32, newPwd string) error {
	err := db.Table("user_info_tb").Where("user_id = ?", userId).Update("user_pwd", newPwd).Error
	return err
}

//更新认证信息
func UpdateUserAuth(userId uint32, userCard, realName string) error {
	data := UserInfo{}
	data.UserCard = userCard
	data.RealName = realName
	err := db.Table("user_info_tb").Where("user_id = ?", userId).Update(&data).Error

	return err
}

func UpdateOrderInfo(userId uint32, usersName, sex, selfDesc, userHead, contactQQ string) error {
	data := UserInfo{}
	data.UserName = usersName
	data.Sex = sex
	data.SelfDesc = selfDesc
	data.UserHead = userHead
	data.ContactQQ = contactQQ
	err := db.Table("user_info_tb").Where("user_id = ?", userId).Update(&data).Error

	return err
}

//
func GetUserList(userAcc string, pageFrom, pageSize int32) (dataList []UserInfo, totalNum int32, err error) {
	thisDb := db.Table("user_info_tb")
	if userAcc != "" {
		thisDb = thisDb.Where("user_acc = ?", userAcc)
	}
	err = thisDb.Count(&totalNum).Error
	if err != nil {
		return dataList, totalNum, err
	}

	err = thisDb.Offset(pageFrom).Limit(pageSize).Find(&dataList).Error

	return dataList, totalNum, err
}

func UpdateUserType(userId uint32) error {
	return db.Table("user_info_tb").Where("user_id = ?", userId).Update("user_type", USERTYPE_BOOSTER).Error
}