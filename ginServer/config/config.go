package config

import (
	"encoding/json"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
)

var ServerConf = new(ServerConfig)
var MysqlConf = new(MysqlConfig)
var LoggerConf = new(LoggerConfig)

//初始化配置
func Init() error {
	//log配置
	log.SetFlags(log.Llongfile | log.Ldate)

	//读取配置
	if cfg, err := ini.Load("ginServer/conf/config.ini"); err != nil {
		log.Println("ginServer/conf/config.ini不存在")
		return err
	} else {
		//server配置
		if err = cfg.Section("server").MapTo(ServerConf); err != nil {
			return err
		}
		//mysql配置
		if err = cfg.Section("mysql").MapTo(MysqlConf); err != nil {
			return err
		}
	}

	//日志配置
	if b, err := ioutil.ReadFile("ginServer/conf/logger.json"); err != nil {
		return err
	} else {
		if err := json.Unmarshal(b, LoggerConf); err != nil {
			return err
		} else {
			log.Println(LoggerConf)
		}
	}
	return nil
}
