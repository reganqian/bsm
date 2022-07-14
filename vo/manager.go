package vo


import (
	"bsm/models"
)

type LoginManager struct {
	ManagerId int32 `json:"managerId"`
	LoginToken string `json:"loginToken"`
	ManagerName string `json:"managerName"`
	LoginTime int64 `json:"loginTime"`
}



type LoginManagerReply struct {
	BaseReply
	Data *LoginManager `json:"data"`
}


type ManagerListReply struct {
	BaseReply
	DataList []models.ManagerInfo `json:"dataList"`
}
