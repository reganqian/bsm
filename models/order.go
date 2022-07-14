package models

import (
	"common/utils"
	. "common/static"
	"time"
)

type OrderInfo struct {
	OrderId string `gorm:"primary_key" json:"orderId"`
	UserId uint32 `json:"userId"`
	GameId int32 `json:"gameId"`
	PlatId int32 `json:"platId"`
	ServerId int32 `json:"serverId"`
	GameAcc string `json:"gameAcc"`	
	GamePwd string `json:"gamePwd"`
	RoleName string `json:"roleName"`

	OrderLimit int32 `json:"orderLimit"`//时限单位小时
	OrderAmount float64 `json:"orderAmount"`
	OrderBond float64 `json:"orderBond"`
	EfficiencyBond float64 `json:"efficiencyBond"`
	OrderDesc string `json:"orderDesc"`
	OrderReq string `json:"orderReq"`
	OrderStatus string `json:"orderStatus"`
	Remark string `json:"remark"`
	ContactUser string `json:"contactUser"`
	ContactPhone string `json:"contactPhone"`
	ContactQQ string `json:"contactQQ"`
	DeadLine int64  `json:"deadLine"`

	PlayerId uint32 `json:"playerId"`
	StartTime int64 `json:"startTime"`
	FinishTime int64 `json:"finishTime"`
	OverTime int64 `json:"overTime"`
	Score int32 `json:"score"`

	StartLevel int32 `json:"startLevel"`
	StartStar int32 `json:"startStar"`
	EndLevel int32 `json:"endLevel"`
	EndStar int32 `json:"endStar"`
	CreateTime int64 `json:"createTime"`

	OrderType string `json:"orderType"`
}

type OrderDetail struct {
	OrderInfo
	GameName string `json:"gameName"`
	PlatName string `json:"platName"`
	ServerName string `json:"serverName"`

}

func SaveOrderInfo(userId uint32, gameId, platId, serverId, orderLimit int32, 
	orderAmount, orderBond, efficiencyBond float64,
	orderType, gameAcc, gamePwd, roleName, orderDesc, orderReq, contactUser, contactPhone, contactQQ string,
	orderLength, startLevel, startStar, endLevel, endStar int32) error {
	data := OrderInfo{}
	data.OrderId = utils.CreateOrderNo("")
	data.UserId = userId
	data.GameId = gameId
	data.PlatId = platId
	data.ServerId = serverId
	data.GameAcc = gameAcc
	data.GamePwd = gamePwd
	data.RoleName = roleName
	data.OrderLimit = orderLimit
	data.OrderAmount = orderAmount
	data.OrderBond = orderBond
	data.EfficiencyBond = efficiencyBond
	data.OrderDesc = orderDesc
	data.OrderReq = orderReq
	data.OrderStatus = ORDER_STATUS_CREATE
	data.ContactUser = contactUser
	data.ContactPhone = contactPhone
	data.ContactQQ = contactQQ
	nowTime := time.Now().Unix()
	data.CreateTime = nowTime
	data.DeadLine = nowTime + int64(orderLength * 24 * 60 * 60)
	data.StartLevel = startLevel
	data.StartStar = startStar
	data.EndLevel = endLevel
	data.EndStar = endStar
	data.OrderType  = orderType
	err := db.Table("order_info_tb").Create(&data).Error
	return err
}

//获取订单列表
func GetOrderList(userId uint32, gameId int32, orderDesc, orderType string, smallAmount, bigAmount float64,
	 pageFrom, pageSize int32 ) (dataList []OrderDetail, totalNum int32, err error) {
	
	thisDb := db.Table("order_info_tb").Where("order_info_tb.game_id = ?", gameId)
	thisDb = thisDb.Joins("left join game_info_tb on order_info_tb.game_id = game_info_tb.game_id")
	thisDb = thisDb.Joins("left join game_plat_tb on game_plat_tb.plat_id = order_info_tb.plat_id")
	thisDb = thisDb.Joins("left join game_server_tb on game_server_tb.server_id = order_info_tb.server_id")
	if userId != 0 {
		thisDb = thisDb.Where("user_id = ?", userId)
	}
	if orderDesc != "" {
		thisDb = thisDb.Where("order_desc = ? or order_req = ?", orderDesc, orderDesc)
	}
	if orderType != "" {
		thisDb = thisDb.Where("order_type = ? ", orderType)
	}
	if smallAmount != 0 {
		thisDb.Where("order_amount >= ?", smallAmount)
	}
	if bigAmount != 0 {
		thisDb.Where("order_amount <= ?", bigAmount)
	}
	err = thisDb.Count(&totalNum).Error
	if err != nil {
		return dataList, totalNum, err
	}

	err = thisDb.Select("order_info_tb.*,game_server_tb.server_name, game_info_tb.game_name, game_plat_tb.plat_name").Offset(pageFrom).Limit(pageSize).Find(&dataList).Error
	return dataList, totalNum, err
}


func GetOrderById(orderId string) (data OrderInfo, err error) {
	err = db.Table("order_info_tb").Where("order_id = ?", orderId).First(&data).Error
	return data, err
}

func UserAcceptOrder(userId uint32, orderId string) error {
	data := OrderInfo{}
	data.PlayerId = userId
	data.StartTime = time.Now().Unix()
	data.OrderStatus = ORDER_STATUS_WORKING//修改为执行中
	err := db.Table("order_info_tb").Where("order_id = ?", orderId).Update(&data).Error
	return err
}