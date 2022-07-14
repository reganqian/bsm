package models

import (
	"common/utils"
)

type UserAuth struct {
	AuthId string `gorm:"primary_key" json:"authId"`
	UserId uint32 `json:"userId"`
	AuthStatus string `json:"authStatus"`
	UserCard string `json:"userCard"`
	RealName string `json:"realName"`
	CreateTime int64 `json:"createTime"`
}


//新增认证记录
func AddUserAuth() error  {
	data := UserAuth{}
	data.AuthId = utils.GetIdStr()

	return nil
}

func GetUserAuthList(pageFrom, pageSize int32) (dataList []UserAuth, totalNum int32, err error) {
	thisDb := db.Table("user_auth_tb")
	err = thisDb.Count(&totalNum).Error
	if err != nil {
		return dataList, totalNum, err
	}
	
	err = thisDb.Offset(pageFrom).Limit(pageSize).Find(&dataList).Error
	return dataList, totalNum, err
}
