package vo

import (
	"bsm/models"
)

//用户列表
type UserListReply struct {
	BaseReply
	TotalNum int32 `json:"totalNum"`
	DataList []models.UserInfo `json:"dataList"`
}

//用户认证列表
type UserAuthListReply struct {
	BaseReply
	TotalNum int32 `json:"totalNum"`
	DataList []models.UserAuth `json:"dataList"`
}

//用户游戏认证列表
type UserGameAuthRecordsReply struct {
	BaseReply
	TotalNum int32 `json:"totalNum"`
	DataList []models.UserGameAuthDetail `json:"dataList"`
}