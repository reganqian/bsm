package service

import (
	"bsm/vo"
	. "common/static"
	"common/utils"
	"bsm/models"
	"time"
	"fmt"
	"encoding/json"
	"bsm/redis"
)

//保存登录信息到redis
func SaveLoginInfoToRedis(loginInfo vo.LoginManager) {
	//保存到redis
	b, _ := json.Marshal(loginInfo)
	redis.RedisSetExDefault(loginInfo.LoginToken, string(b))
}


//密码登录
func ManagerLogin(managerAcc, managerPwd string) vo.LoginManagerReply {
	reply := vo.LoginManagerReply{}

	//1:根据手机号查询账号
	data, err := models.QueryByAcc(managerAcc)
	if err != nil {
		reply.Failed(DBERROR, "查询错误" + err.Error())
		return reply
	}

	//2:判断密码是否正确
	if managerPwd != data.ManagerPwd {
		//不正确， 
		reply.Failed(PWDERROR, "密码错误")
		return reply
	}

	//保存登录信息
	loginInfo := vo.LoginManager{}
	tokensalt := utils.GetIdStr()
	loginInfo.ManagerId = data.ManagerId
	tokenStr := utils.CreateMd5String(managerAcc, tokensalt)
	loginInfo.LoginToken = tokenStr
	loginInfo.LoginTime = time.Now().Unix()

	
	//保存
	SaveLoginInfoToRedis(loginInfo)

	//TODO 填写登录日志
	reply.Success()
	reply.Data = &loginInfo

	return reply
}

func CheckToken(token string) vo.LoginManagerReply {
	loginReply := DoAuth(token)
	//更新时间
	return loginReply
}

//执行验证流程
func  DoAuth(token string) (reply vo.LoginManagerReply) {
	v, err1 := redis.RedisGet(token);
	if v == nil || err1 != nil {
		reply.ResCode = FAILED
		reply.ResDesc =  "user is not exist!"
		return reply
	}

	var b []uint8
    b = v.([]uint8)
	var auth  vo.LoginManager
    if err := json.Unmarshal(b, &auth); err == nil {

		// utils.RedisSetExDefault(token, string(b))
    } else {
		reply.ResCode = FAILED
		reply.ResDesc =  "data fomat error!"
		return reply
	}
	
	reply.ResCode = SUCCESS
	reply.ResDesc =  DEFAULT_SUCCESS_DESC
	reply.Data = &auth
	fmt.Println(auth)
	return reply
	// return SUCCESS, "success", auth.Data.UserId
}



func CreateManager(managerAcc, managerPwd, managerRole, managerName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.CreateManager(managerAcc, managerPwd, managerRole, managerName)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.Success()
	return reply
}

func UpdateManager(managerId int32, managerAcc, managerPwd, managerRole, managerName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.UpdateManager(managerId, managerAcc, managerPwd, managerRole, managerName)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.Success()
	return reply
}

func GetManagerList() vo.ManagerListReply {
	reply := vo.ManagerListReply{}
	dataList, err := models.GetManagerList()
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.DataList = dataList
	reply.Success()
	return reply
}