package service

import (
	"bsm/vo"
	. "common/static"
	"bsm/models"
)

func GetGameList() vo.GameListReply {
	reply := vo.GameListReply{}
	dataList, err := models.GetGameList()
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.DataList = dataList

	reply.Success()
	return reply 
}

func AddGame(gameName, gameLogo string, season int32) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.AddGame(gameName, gameLogo, season)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.Success()
	return reply
}

func UpdateGame(gameId int32, gameName, gameLogo string, season int32) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.UpdateGame(gameId, gameName, gameLogo, season)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.Success()
	return reply
}



func GetPlatList(gameId int32) vo.PlatListReply {
	reply := vo.PlatListReply{}
	dataList, err := models.GetGetAllPlats(gameId)

	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	resList := []*vo.PlatDetail{}
	gameInfo, err := models.GetGameById(gameId)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	for _, data := range dataList {
		res := vo.PlatDetail{}
		res.GamePlat = data
		res.GameName = gameInfo.GameName
		resList = append(resList, &res)
	}

	reply.DataList = resList

	reply.Success()
	return reply 
}

func AddGamePlat(gameId int32, platName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.AddGamePlat(gameId, platName)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	
	reply.Success()
	return reply
}

func UpdatePlat(platId, gameId int32, platName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.UpdateGamePlat(platId, gameId, platName)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	
	reply.Success()
	return reply
}


func GetServerList(gameId, platId int32) vo.ServerListReply {
	reply := vo.ServerListReply{}

	dataList, err := models.GetServerList(gameId, platId)

	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	resList := []*vo.ServerDetail{}
	gameInfo, err := models.GetGameById(gameId)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	platInfo, err := models.GetPlatById(platId)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	for _, data := range dataList {
		res := vo.ServerDetail{}
		res.GameServer = data
		res.GameName = gameInfo.GameName
		res.PlatName = platInfo.PlatName
		resList = append(resList, &res)
	}

	reply.DataList = resList

	reply.Success()

	return reply 
}

func AddServer(gameId, platId, serverNo int32, serverName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.AddGameServer(gameId, platId, serverNo, serverName) 
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	
	reply.Success()
	return reply
}

func UpdateServer(serverId, gameId, platId, serverNo int32, serverName string) vo.BaseReply {
	reply := vo.BaseReply{}
	err := models.UpdateGameServer(serverId, gameId, platId, serverNo, serverName)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	
	reply.Success()
	return reply
}

func GetAllLevels(gameId int32) vo.LevelsReply {
	reply := vo.LevelsReply{}
	dataList, err := models.GetAllLevels(gameId)
	if err != nil {
		reply.Failed(DBERROR, err.Error())
		return reply
	}
	reply.DataList = dataList
	reply.Success()
	return reply
}


