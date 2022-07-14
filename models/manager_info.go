package models

import (
	"time"
)


type ManagerInfo struct {
	ManagerId int32 `gorm:"primary_key" json:"managerId"`
	ManagerAcc string `json:"managerAcc"`
	ManagerPwd string `json:"managerPwd"`
	ManagerRole string `json:"managerRole"`
	ManagerName string `json:"managerName"`
	CreateTime int64 `json:"createTime"`
}


func CreateManager(managerAcc, managerPwd, managerRole, managerName string) error {
	data := ManagerInfo{}
	data.ManagerAcc = managerAcc
	data.ManagerPwd = managerPwd
	data.ManagerRole = managerRole
	data.ManagerName = managerName
	data.CreateTime = time.Now().Unix()

	return db.Table("manager_info_tb").Create(&data).Error
}


func UpdateManager(managerId int32, managerAcc, managerPwd, managerRole, managerName string) error {
	data := ManagerInfo{}
	data.ManagerAcc = managerAcc
	data.ManagerPwd = managerPwd
	data.ManagerRole = managerRole
	data.ManagerName = managerName

	return db.Table("manager_info_tb").Where("manager_id = ?", managerId).Update(&data).Error
}

func QueryByAcc(managerAcc string) (data ManagerInfo, err error) {
	err = db.Table("manager_info_tb").Where("manager_acc = ?", managerAcc).First(&data).Error

	return data, err
}

//查询管理员列表
func GetManagerList() (dataList []ManagerInfo, err error) {
	err = db.Table("manager_info_tb").Find(&dataList).Error
	return dataList, err
}