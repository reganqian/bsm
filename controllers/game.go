package controllers

import (
	"github.com/astaxie/beego"
	"bsm/service"
	// . "common/static"
)


type GameController struct {
	beego.Controller
}



func (s *GameController) GameList()  {
	reply := service.GetGameList()
	s.Data["json"] = reply
	s.ServeJSON()	
}


func (s *GameController) PlatList()  {
	gameId, _ := s.GetInt("gameId")
	reply := service.GetPlatList(int32(gameId))
	s.Data["json"] = reply
	s.ServeJSON()	
}

func (s *GameController) ServerList()  {
	gameId, _ := s.GetInt("gameId")
	platId, _ := s.GetInt("platId")
	reply := service.GetServerList(int32(gameId) , int32(platId))
	s.Data["json"] = reply
	s.ServeJSON()	
}


func (s *GameController) AddGame()  {
	gameName := s.GetString("gameName")
	gameLogo := s.GetString("gameLogo")
	season, _ := s.GetInt("season")
	reply := service.AddGame(gameName, gameLogo, int32(season))
	s.Data["json"] = reply
	s.ServeJSON()
}

func (s *GameController) UpdateGame()  {
	gameName := s.GetString("gameName")
	gameLogo := s.GetString("gameLogo")
	gameId, _ := s.GetInt("gameId")
	season, _ := s.GetInt("season")
	reply := service.UpdateGame(int32(gameId), gameName, gameLogo, int32(season))
	s.Data["json"] = reply
	s.ServeJSON()
}



func (s *GameController) AddPlat()  {
	gameId, _ := s.GetInt("gameId")
	platName := s.GetString("platName")

	reply := service.AddGamePlat(int32(gameId), platName)
	s.Data["json"] = reply
	s.ServeJSON()
}

func (s *GameController) UpdatePlat()  {
	platName := s.GetString("platName")
	platId, _ := s.GetInt("platId")
	gameId, _ := s.GetInt("gameId")
	reply := service.UpdatePlat(int32(platId), int32(gameId), platName)
	s.Data["json"] = reply
	s.ServeJSON()
}



func (s *GameController) AddServer()  {
	serverName := s.GetString("serverName")
	
	gameId, _ := s.GetInt("gameId")
	platId, _ := s.GetInt("platId")
	serverNo, _ := s.GetInt("serverNo")

	reply := service.AddServer(int32(gameId), int32(platId), int32(serverNo), serverName)
	s.Data["json"] = reply
	s.ServeJSON()
}

func (s *GameController) UpdateServer()  {
	serverName := s.GetString("serverName")
	serverId, _ := s.GetInt("serverId")
	gameId, _ := s.GetInt("gameId")
	platId, _ := s.GetInt("platId")
	serverNo, _ := s.GetInt("serverNo")
	reply := service.UpdateServer(int32(serverId), int32(gameId), int32(platId), int32(serverNo), serverName)
	s.Data["json"] = reply
	s.ServeJSON()
}

//获取游戏段位信息
func (s *GameController) GetAllLevels()  {
	gameId, _ := s.GetInt("gameId")
	reply := service.GetAllLevels(int32(gameId))
	s.Data["json"] = reply
	s.ServeJSON()	
}
