package models

import (
	"time"
)

type GameInfo struct {
	GameId int32 `gorm:"primary_key" json:"gameId"`
	GameName string `json:"gameName"`
	GameLogo string `json:"gameLogo"`
	CreateTime int64 `json:"createTime"`
	Season int32 `json:"season"`//赛季
}

type GamePlat struct {
	PlatId int32 `gorm:"primary_key" json:"platId"`
	GameId int32 `json:"gameId"`
	PlatName string `json:"platName"`
}

type GameServer struct {
	ServerId int32 `gorm:"serverId" json:"serverId"`
	GameId int32 `json:"gameId"`
	PlatId int32 `json:"platId"`
	ServerNo int32 `json:"serverNo"`
	ServerName string `json:"serverName"`
	CreateTime int64 `json:"createTime"`
}

//查询所有游戏列表
func GetGameList() (dataList []GameInfo, err error) {
	err = db.Table("game_info_tb").Find(&dataList).Error
	return dataList, err
}

func GetGameById(gameId int32) (data GameInfo, err error) {
	err = db.Table("game_info_tb").Where("game_id= ?", gameId).First(&data).Error
	return data, err
}

//平台列表
func GetGetAllPlats(gameId int32) (dataList []GamePlat, err error) {
	err = db.Table("game_plat_tb").Where("game_id = ?", gameId).Find(&dataList).Error
	return dataList, err	
}


func GetPlatById(platId int32) (data GamePlat, err error) {
	err = db.Table("game_plat_tb").Where("plat_id  = ?", platId).First(&data).Error
	return data, err	
}


//查询服务器列表
func GetServerList(gameId, platId int32) (dataList []GameServer, err error) {
	err = db.Table("game_server_tb").Where("game_id = ? and plat_id = ?", gameId, platId).Find(&dataList).Error
	return dataList, err	
}


func GetServerById (serverId int32) (data GameServer, err error) {
	err = db.Table("game_server_tb").Where("server_id = ?", serverId).First(&data).Error
	return data, err
}

func AddGame(gameName, gameLogo string, season int32) error {
	data := GameInfo{}
	data.GameName = gameName
	data.GameLogo = gameLogo
	data.Season = season
	data.CreateTime = time.Now().Unix()
	return db.Table("game_info_tb").Create(&data).Error
}

func UpdateGame(gameId int32, gameName, gameLogo string, season int32) error {
	data := GameInfo{}
	data.GameName = gameName
	data.GameLogo = gameLogo
	data.Season = season
	return db.Table("game_info_tb").Where("game_id = ?", gameId).Update(&data).Error
}

func AddGamePlat(gameId int32, platName string) error {
	data := GamePlat{}
	data.GameId = gameId
	data.PlatName = platName
	return db.Table("game_plat_tb").Create(&data).Error
}

func UpdateGamePlat(platId, gameId int32, platName string) error {
	data := GamePlat{}
	data.GameId = gameId
	data.PlatName = platName
	return db.Table("game_plat_tb").Where("plat_id = ?", platId).Update(&data).Error
}


func AddGameServer(gameId, platId, serverNo int32, serverName string) error {
	data := GameServer{}
	data.GameId = gameId
	data.PlatId = platId
	data.ServerNo = serverNo
	data.ServerName = serverName
	data.CreateTime = time.Now().Unix()
	return db.Table("game_server_tb").Create(&data).Error
}



func UpdateGameServer(serverId, gameId, platId, serverNo int32, serverName string) error {
	data := GameServer{}
	data.GameId = gameId
	data.PlatId = platId
	data.ServerNo = serverNo
	data.ServerName = serverName
	data.CreateTime = time.Now().Unix()
	return db.Table("game_server_tb").Where("server_id = ?", serverId).Update(&data).Error
}


