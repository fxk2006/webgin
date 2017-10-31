package global

import (
	"webgin/util"
	"github.com/robfig/config"
	"log"
	"strings"
)

const configName = "conf.ini"
const logName = "log.log"

var readConfig = func(name string) (*config.Config) {
	conf, err := config.ReadDefault(configName)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

var Config = readConfig(configName)

var getLogLevel = func(conf *config.Config) (byte) {
	level, _ := conf.String("log", "level")
	//忽略大小写比较
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

/*var Debug = func(v ...interface{}) {
	glog.Debug(v...)
}
var Info = func(v ...interface{}) {
	glog.Info(v...)
}
var Warning = func(v ...interface{}) {
	glog.Warning(v...)
}*/
