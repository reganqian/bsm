// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	control "bsm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/dlm",

	beego.NSNamespace("/manager", 
		beego.NSRouter("/create", &control.ManagerController{}, "post:CreateManager"), 

		beego.NSRouter("/update", &control.ManagerController{}, "post:UpdateManager"), 

		beego.NSRouter("/list", &control.ManagerController{}, "post:GetManagerList"), 

		// beego.NSRouter("/accept", &control.ManagerController{}, "post:AcceptOrder"), 
		//
	),

	beego.NSNamespace("/user", 
		beego.NSRouter("/userList", &control.UserController{}, "post:GetUserList"), 

		beego.NSRouter("/authList", &control.UserController{}, "post:GetUserAuthList"), 
		beego.NSRouter("/gameAuthRecords", &control.UserController{}, "post:GetUserGameAuthRecords"), 
		beego.NSRouter("/authUserGame", &control.UserController{}, "post:AuthUserGame"), 
		
		
	),

	beego.NSNamespace("/order", 
		beego.NSRouter("/list", &control.OrderController{}, "post:GetOrderList"), 
		
		
	),
	//

	beego.NSNamespace("/game", 
		beego.NSRouter("/gameList", &control.GameController{}, "post:GameList"), 
		beego.NSRouter("/platList", &control.GameController{}, "post:PlatList"), 
		beego.NSRouter("/serverList", &control.GameController{}, "post:ServerList"), 


		beego.NSRouter("/addGame", &control.GameController{}, "post:AddGame"), 
		beego.NSRouter("/updateGame", &control.GameController{}, "post:UpdateGame"), 
		beego.NSRouter("/addPlat", &control.GameController{}, "post:AddPlat"), 
		beego.NSRouter("/updatePlat", &control.GameController{}, "post:UpdatePlat"), 
		beego.NSRouter("/addServer", &control.GameController{}, "post:AddServer"), 
		beego.NSRouter("/updateServer", &control.GameController{}, "post:UpdateServer"), 

		beego.NSRouter("/allLevels", &control.GameController{}, "post:GetAllLevels"), 

		// beego.NSRouter("/accept", &control.ManagerController{}, "post:AcceptOrder"), 
		//
	),

	

	beego.NSRouter("/login", &control.ManagerController{}, "post:ManagerLogin"), //登录
	
	
	)

	beego.AddNamespace(ns)
}
