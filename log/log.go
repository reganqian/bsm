package log

import (
	
	"github.com/sirupsen/logrus"
	"fmt"
	graylog "github.com/gemnasium/logrus-graylog-hook"
	"github.com/astaxie/beego/config"
)

var Log = logrus.New()

func init() {

	log_open := "No"
	iniconf, err1 := config.NewConfig("ini", "conf/config.ini")
	if err1 != nil {
		Log.Error(err1.Error())
	} else {
	// 2. 通过对象获取数据
		log_host := iniconf.String("server::log_addr")
		log_port := iniconf.String("server::log_port")
		log_open = iniconf.String("sys::log_open")
		if log_open == "YES" {
			graylog_addr := fmt.Sprintf("%s:%s", log_host, log_port)
			hook := graylog.NewGraylogHook(graylog_addr, map[string]interface{}{"Project": "52web"})
			Log.AddHook(hook)
		}
	}
	

}