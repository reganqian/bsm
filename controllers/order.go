package controllers



import (
	"github.com/astaxie/beego"
	"bsm/service"
	// . "common/static"
)


type OrderController struct {
	beego.Controller
}



func (s *OrderController) GetOrderList()  {
	orderDesc := s.GetString("orderDesc")//非必填
	orderType := s.GetString("orderType")//非必填
	

	userId, _ := s.GetInt("userId")
	gameId, _ := s.GetInt("gameId")
	pageNo, _ := s.GetInt("pageNo")
	pageSize, _ := s.GetInt("pageSize")

	smallAmount, _ := s.GetFloat("smallAmount")
	bigAmount, _ := s.GetFloat("bigAmount")


	reply := service.GetOrderList(uint32(userId), int32(gameId), orderDesc, orderType, smallAmount, bigAmount,
	int32(pageNo), int32(pageSize))
	s.Data["json"] = reply
	s.ServeJSON()	
}