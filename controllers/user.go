package controllers


import (
	"github.com/astaxie/beego"
	"bsm/service"
	// . "common/static"
)


type UserController struct {
	beego.Controller
}



func (s *UserController) GetUserList()  {
	userAcc := s.GetString("userAcc")//非必填
	pageNo, _ := s.GetInt("pageNo")
	pageSize, _ := s.GetInt("pageSize")
	reply := service.UserList(userAcc, int32(pageNo), int32(pageSize))
	s.Data["json"] = reply
	s.ServeJSON()	
}


func (s *UserController) GetUserAuthList()  {
	pageNo, _ := s.GetInt("pageNo")
	pageSize, _ := s.GetInt("pageSize")
	reply := service.GetUserAuthList(int32(pageNo), int32(pageSize))
	s.Data["json"] = reply
	s.ServeJSON()	
}


func (s *UserController) GetUserGameAuthRecords()  {
	authStatus := s.GetString("authStatus")//非必填
	userId,  _ := s.GetInt("userId")
	pageNo, _ := s.GetInt("pageNo")
	pageSize, _ := s.GetInt("pageSize")
	reply := service.GetUserGameAuthRecords(uint32(userId), authStatus, int32(pageNo), int32(pageSize))
	s.Data["json"] = reply
	s.ServeJSON()	
}



func (s *UserController) AuthUserGame()  {
	authStatus := s.GetString("authStatus")
	remark := s.GetString("remark")
	recordId,  _ := s.GetInt("recordId")

	reply := service.AuthUserGame(int32(recordId), authStatus, remark) 
	s.Data["json"] = reply
	s.ServeJSON()	
}


