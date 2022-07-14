package vo

import (
	"bsm/models"
)

type GameListReply struct {
	BaseReply
	DataList []models.GameInfo `json:"dataList"`
}

type PlatListReply struct {
	BaseReply
	DataList []*PlatDetail `json:"dataList"`
}

type PlatDetail struct {
	models.GamePlat
	GameName  string `json:"gameName"`
}

type ServerListReply struct {
	BaseReply
	DataList []*ServerDetail `json:"dataList"`
}

type ServerDetail struct {
	models.GameServer
	GameName  string `json:"gameName"`
	PlatName string `json:"platName"`
}
	
type LevelsReply struct {
	BaseReply
	DataList []models.GameLevel `json:"dataList"`
}