package service

import (
	"bsm/vo"
	"bsm/models"
	. "common/static"
)


func GetOrderList(userId uint32, gameId int32, orderDesc, orderType string, smallAmount, bigAmount float64,
	pageNo, pageSize int32 ) vo.OrderListReply {
	pageFrom := int32(0)
	reply := vo.OrderListReply{}
	pageNo, pageSize, pageFrom = CheckPageReq(pageNo, pageSize, pageFrom)
	dataList, totalNum, err := models.GetOrderList(userId, gameId, orderDesc, orderType, smallAmount, bigAmount, pageFrom, pageSize)
	if err != nil {
		reply.Failed(DBERROR, "查询订单数据异常：" + err.Error())
		return reply
	}

	reply.DataList = dataList
	reply.TotalNum = totalNum
	reply.Success()
	return reply
}
