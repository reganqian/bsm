package models

import (
	"common/utils"
	. "common/static"
	"time"
)

type PhoneCode struct {
	CodeId string `gorm:"primary_key" json:"codeId"`
	UserId uint32 `json:"userId"`
	PhoneNum string `json:"phoneNum"`
	PhoneCode string `json:"PhoneCode"`
	CodeType string `json:"codeType"`
	CreateTime int64  `json:"createTime"`
	CodeStatus string `json:"codeStatus"`

}

//保存记录
func SavePhoneCode(phoneNum, phoneCode, codeType string) error {
	data := PhoneCode{}
	data.CodeId = utils.GetIdStr()
	data.PhoneNum = phoneNum
	data.PhoneCode = phoneCode
	data.CodeType = codeType
	data.CreateTime = time.Now().Unix()
	data.CodeStatus = STATUS_NOT_USE
	err := db.Table("phone_code_tb").Create(&data).Error
	return err
}

//根据手机号和类型查询出数据
func QueryCodeByPhoneAndType(phoneNum, codeType string) (dataList []PhoneCode, err error) {
	err = db.Table("phone_code_tb").Where("phone_num = ? and code_type = ?", phoneNum, codeType).Find(&dataList).Error

	return dataList, err
}

//使用验证码
func UseCode(codeId string) error {
	err := db.Table("phone_code_tb").Where("code_id = ?", codeId).Update("code_status", STATUS_USED).Error

	return err
}

//验证码过期
func TimeOutCode(codeId string) error {
	err := db.Table("phone_code_tb").Where("code_id = ?", codeId).Update("code_status", STATUS_TIMEOUT).Error

	return err
}

