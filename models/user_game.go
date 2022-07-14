package models

import (
	"time"
)

type UserGame struct {
	UgdId uint32 `gorm:"primary_key" json:"UgdId"`
	UserId uint32 `json:"userId"`
	GameId int32  `json:"gameId"`
	PlatId int32 `json:"platformId"`
	ServerId int32 `json:"serverId"`
	AuthStatus string `json:"authStatus"`
	GameLevel int32 `json:"gameLevel"`
	CreateTime int64 `json:"createTime"`
	RoleName string `json:"roleName"`
	LevelStar int32 `json:"levelStar"`
	AuthTime int64 `json:"authTime"`
}

func AddUserGame(userId uint32, gameId, platId, serverId, gameLevel, levelStar int32, roleName string, authTime int64) error {
	data := UserGame{}
	data.UserId = userId
	data.GameId = gameId
	data.PlatId = platId
	data.ServerId = serverId
	data.CreateTime = time.Now().Unix()
	data.RoleName = roleName
	data.GameLevel = gameLevel
	data.LevelStar = levelStar
	data.AuthTime = authTime

	return db.Table("user_game_tb").Create(&data).Error
}

func GetGameInfoByUser(userId uint32) (dataList []UserGame, err error) {
	err = db.Table("user_game_tb").Where("user_id = ?", userId).Find(&dataList).Error
	return dataList, err 
}