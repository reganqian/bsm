package models


import (
	"time"
	. "common/static"
)

type UserGameAuthRecord struct {
	RecordId int32 `gorm:"primary_key" json:"recordId"`
	UserId uint32 `json:"userId"`
	GameId int32 `json:"gameId"`
	PlatId int32 `json:"platId"`
	RoleName string `json:"roleName"`
	CreateTime int64 `json:"createTime"`
	AuthDesc string `json:"authDesc"`
	Remark string `json:"remark"`
	AuthStatus string `json:"authStatus"`
	GameLevel int32 `json:"gameLevel"`
	LevelStar int32 `json:"levelStar"`
	Season	int32 `json:"season"`
}

func SaveUserGameAuthRecord(userId uint32, gameId, platId, gameLevel, levelStar, season int32, 
	roleName, authDesc string) error {
	data := UserGameAuthRecord{}
	data.UserId = userId
	data.GameId = gameId
	data.PlatId = platId
	data.RoleName = roleName
	data.CreateTime = time.Now().Unix()
	data.AuthDesc = authDesc
	data.AuthStatus = AUTH_STATUS_APPLY
	data.GameLevel = gameLevel
	data.LevelStar = levelStar
	data.Season = season

	err := db.Table("user_game_auth_record_tb").Create(&data).Error

	return err
}

type UserGameAuthDetail struct {
	UserGameAuthRecord
	GameName string `json:"gameName"`
	LevelName string `json:"levelName"`
	PlatName string `json:"platName"`
}

func GetAuthRecordById(recordId int32) (data UserGameAuthRecord, err error) {
	err = db.Table("user_game_auth_record_tb").Where("record_id = ?", recordId).First(&data).Error
	return data, err
}

func UpdateAuthStatus(recordId int32, authStatus, remark string) error  {
	data := UserGameAuthRecord{}
	data.AuthStatus = authStatus
	if remark != "" {
		data.Remark = remark
	}
	err := db.Table("user_game_auth_record_tb").Where("record_id = ?", recordId).Update(&data).Error
	return err
}

func GetUserGameAuthRecordList(userId uint32, authStatus string, pageSize, pageFrom int32) (dataList []UserGameAuthDetail, totalNum int32, err error) {
	thisDb := db.Table("user_game_auth_record_tb").Joins("left join game_info_tb on user_game_auth_record_tb.game_id = game_info_tb.game_id")
	thisDb = thisDb.Joins("left join game_level_tb on game_level_tb.level_id = user_game_auth_record_tb.game_level")
	thisDb = thisDb.Joins("left join game_plat_tb on game_plat_tb.plat_id = user_game_auth_record_tb.plat_id")
	if authStatus != "" {
		thisDb = thisDb.Where("auth_status = ?", authStatus)
	}
	if userId != 0 {
		thisDb = thisDb.Where("user_id = ?", userId)
	}
	err = thisDb.Count(&totalNum).Error
	if err != nil {
		return dataList, totalNum, err
	}

	err = thisDb.Select("user_game_auth_record_tb.*, game_info_tb.game_name, game_level_tb.level_name, game_plat_tb.plat_name").Offset(pageFrom).Limit(pageSize).Find(&dataList).Error
	return dataList, totalNum, err
}

