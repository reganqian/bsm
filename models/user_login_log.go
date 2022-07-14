package models

import (
	"time"
	"common/utils"
	. "common/static"
	// "errors"
)

type UserLoginLog struct {
	LogId string `gorm:"primary_key" json:"logId"`
	UserId uint32 `json:"userId"`
	LoginIp string `json:"loginIp"`
	LoginTime int64 `json:"loginTime"`
	LoginStatus string `json:"loginStatus"`
	LoginServer string `json:"loginServer"`
	LoginPhone string `json:"loginPhone"`
}

//保存登录记录
func SaveLoginLog(userId uint32, loginIp, loginStatus, loginServer, loginPhone string) error {

	data := UserLoginLog{}
	data.LogId = utils.GetIdStr()
	data.UserId = userId
	data.LoginIp = loginIp
	data.LoginTime = time.Now().Unix()
	data.LoginStatus = loginStatus
	data.LoginServer = loginServer
	data.LoginPhone = loginPhone
	err := db.Table("user_login_log_tb").Create(&data).Error
	if loginStatus == STATUS_SUCCESS {
		if err == nil {
			//修改用户信息
			userInfo := UserInfo{}
			userInfo.LastLoginIp = loginIp
			userInfo.LastLoginServer = loginServer
			userInfo.LastLoginTime = data.LoginTime
			err = db.Table("user_info_tb").Where("user_id = ?", userId).Update(&userInfo).Error
			if err != nil {
				return err
			}

		}	
	}
	
	return err
}