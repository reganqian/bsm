package models

import (
	// log "github.com/sirupsen/logrus"
	. "bsm/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "time"
	"github.com/astaxie/beego/config"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) Begin() {
	db.db = db.db.Begin()
}

func (db *DB) Rollback() {
	db.db = db.db.Rollback()
}

func (db *DB) Commit() {
	db.db = db.db.Commit()
}

var (
	db *gorm.DB
)

func NewDB() *DB {
	return &DB{db: db}
}

func init() {
	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
	   Log.Error(err1.Error())
	}
	
	// 2. 通过对象获取数据
	dbUrl := iniconf.String("mysql::db_url")
	dbName := iniconf.String("mysql::db_name")
	dbPwd := iniconf.String("mysql::db_pwd")
	
	var err error
	db, err = gorm.Open("mysql", dbName + ":" + dbPwd + "@" + dbUrl)
	if err != nil {
		panic("failed to connect database")
	}
	
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return defaultTableName + "_tb";
	}
	// 自动同步表结构
	// db.SetLogger(logs.GetLogger("orm"))
	db.LogMode(true)
 
}

type Model struct {
	ID        int64       `gorm:"primary_key" json:"id"`
}


