package controllers



import (
	"github.com/astaxie/beego"
	"bsm/service"
	// . "common/static"
)


type ManagerController struct {
	beego.Controller
}

//管理员登录
func (s *ManagerController) ManagerLogin()  {
	managerAcc := s.GetString("managerAcc")
	managerPwd := s.GetString("managerPwd")

	reply := service.ManagerLogin(managerAcc, managerPwd)
	s.Data["json"] = reply
	s.ServeJSON()	
}


//查询列表
func (s *ManagerController) GetManagerList()  {

	reply := service.GetManagerList()
	s.Data["json"] = reply
	s.ServeJSON()	
}
//


//添加管理员
func (s *ManagerController) CreateManager()  {
	managerAcc := s.GetString("managerAcc")
	managerPwd := s.GetString("managerPwd")
	managerRole := s.GetString("managerRole")
	managerName := s.GetString("managerName")

	reply := service.CreateManager(managerAcc, managerPwd, managerRole, managerName )
	s.Data["json"] = reply
	s.ServeJSON()	
}
//

func (s *ManagerController) UpdateManager()  {
	managerAcc := s.GetString("managerAcc")
	managerPwd := s.GetString("managerPwd")
	managerRole := s.GetString("managerRole")
	managerName := s.GetString("managerName")

	managerId, _ := s.GetInt("managerId")

	reply := service.UpdateManager(int32(managerId), managerAcc, managerPwd, managerRole, managerName )
	s.Data["json"] = reply
	s.ServeJSON()	
}