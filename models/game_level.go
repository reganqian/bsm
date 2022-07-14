package models


type GameLevel struct {
	LevelId int32 `gorm:"primary_key" json:"levelId"`
	GameId int32 `json:"gameId"`
	LevelName string `json:"levelName"`
	LevelStar int32 `json:"levelStar"`
	DefaultPrice float64 `json:"defaultPrice"`
	Level int32 `json:"level"`
	LevelTitle string `json:"levelTitle"`
	MinStar int32 `json:"minStar"`

}

func GetAllLevels(gameId int32) (dataList []GameLevel, err error) {
	err = db.Table("game_level_tb").Where("game_id = ?", gameId).Order("level").Find(&dataList).Error
	return dataList, err
}


func GetLevelById(levelId int32) (data GameLevel, err error) {
	err = db.Table("game_level_tb").Where("level_id = ?", levelId).First(&data).Error
	return data, err
}