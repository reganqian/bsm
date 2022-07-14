package main

import (
	_ "bsm/routers"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"bsm/service"
	. "common/static"
	// "github.com/astaxie/beego/config"
	
)

func main() {
	
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//权限校验过滤器
	var FilterUserToken = func(ctx *context.Context) { //这里要在头部importcontext这个包否则无法使用
		auth := ctx.Input.Header("Authorization")
		if auth != "" {
			reply := service.CheckToken(auth) 
			
			if reply.ResCode == SUCCESS {
				ctx.Input.SetData("LoginId", reply.Data.ManagerId) //验证成功设定一个可在任何控制器内获取到的全局变量(用户id)，获取方法是u.Ctx.Input.GetData("UserId").(string)			
			} 
		}
		// Log.Info("in the before")
	} 


	//过滤, 如果需要其他的拦截器, 需要添加参数 false
	beego.InsertFilter("/dlm/*/*", beego.BeforeRouter, FilterUserToken)


	
	beego.Run()
}