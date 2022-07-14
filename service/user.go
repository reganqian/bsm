package service

import (
	"bsm/vo"
	"bsm/models"
	. "common/static"
	"time"
)

//用户列表
func UserList(userAcc string, pageNo, pageSize int32) vo.UserListReply {
	pageFrom := int32(0)
	pageNo, pageSize, pageFrom = CheckPageReq(pageNo, pageSize, pageFrom)
	reply := vo.UserListReply{}
	dataList, totalNum, err := models.GetUserList(userAcc, pageFrom, pageSize)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.TotalNum = totalNum
	reply.DataList = dataList
	reply.Success()
	return reply
}



func GetUserAuthList(pageNo, pageSize int32) vo.UserAuthListReply {
	reply := vo.UserAuthListReply{}
	pageFrom := int32(0)
	pageNo, pageSize, pageFrom = CheckPageReq(pageNo, pageSize, pageFrom)
	dataList, totalNum, err := models.GetUserAuthList(pageFrom, pageSize)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.TotalNum = totalNum
	reply.DataList = dataList
	reply.Success()
	return reply
}

//获取游戏列表
func GetUserGameAuthRecords(userId uint32, authStatus string, pageNo, pageSize int32) vo.UserGameAuthRecordsReply {
	reply := vo.UserGameAuthRecordsReply{}

	pageFrom := int32(0)
	pageNo, pageSize, pageFrom = CheckPageReq(pageNo, pageSize, pageFrom)
	dataList, totalNum, err := models.GetUserGameAuthRecordList(userId, authStatus, pageSize, pageFrom)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.TotalNum = totalNum
	reply.DataList = dataList
	reply.Success()
	return reply
}

//
func AuthUserGame(recordId int32, authStatus, remark string) vo.BaseReply {
	reply := vo.BaseReply{}
	authInfo, err := models.GetAuthRecordById(recordId)
	if err != nil {
		reply.Failed(DBERROR, "获取申请记录异常：" + err.Error())
		return reply
	}
	if authStatus == AUTH_STATUS_PASS {//通过
		//通过需要保存用户游戏信息
		err = models.AddUserGame(authInfo.UserId, authInfo.GameId, authInfo.PlatId, 0, authInfo.GameLevel, authInfo.LevelStar, authInfo.RoleName, time.Now().Unix())
		//修改申请状态
		if err != nil {
			reply.Failed(DBERROR, "添加信息失败：" + err.Error())
			return reply
		}
		
	}
	//修改申请状态
	err  = models.UpdateAuthStatus(recordId, authStatus, remark)
	if err != nil {
		reply.Failed(DBERROR, "修改申请状态失败" + err.Error())
		return reply
	}
	err = models.UpdateUserType(authInfo.UserId)
	if err != nil {
		reply.Failed(DBERROR, "修改用户状态失败" + err.Error())
		return reply
	}
	reply.Success()
	return reply
}