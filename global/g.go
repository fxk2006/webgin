package global

import (
	"webgin/util"
	"github.com/robfig/config"
	"log"
	"strings"
)


var Config = readConfig(configName)

var logName, _ = Config.String(LOG, "name")

var readConfig = func(name string) (*config.Config) {
	conf, err := config.ReadDefault(configName)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

var getLogLevel = func(conf *config.Config) (byte) {
	level, _ := conf.String("log", "level")
	//忽略大小写比较
	if strings.EqualFold("error", level) {
		return util.Error
	}
	if strings.EqualFold("debug", level) {
		return util.Debug
	}
	if strings.EqualFold("info", level) {
		return util.Info
	}
	if strings.EqualFold("warning", level) {
		return util.Warning
	} else {
		return util.Info
	}
}
var GLog *util.Log = util.New(getLogLevel(Config), logName)

var LogError= func(err error){
	if err!=nil{
		GLog.Error(err)
	}
}
