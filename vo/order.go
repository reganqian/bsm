package vo

import (
	"bsm/models"
)



type OrderListReply struct {
	BaseReply
	TotalNum int32 `json:"totalNum"`
	DataList []models.OrderDetail `json:"dataList"`
}