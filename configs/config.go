package run

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig(conf *Config, path string, filteType string) {
	var confObj = viper.New()
	confObj.SetConfigType(filteType)
	confObj.AddConfigPath(path)
	//读取配置文件内容
	if err := confObj.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := confObj.Unmarshal(&conf); err != nil {
		panic(err)
	}
}

var Conf Config

func InitConfig() {
	env := os.Getenv("ENV")
	var path string
	if env == "DEBUG" {
		path = "./../role-system/configs/dev"
	} else if env == "PROD" {
		path = "./../role-system/configs/prod"
	} else {
		panic("请设置环境变量: ENV=DEBUG(测试) ENV=PROD(正式)")
	}
	LoadConfig(&Conf, path, "yaml")
}

type Config struct {
	RunMode string
	HTTP    Http
	Casbin  Casbin
	MySQL   MySQL
	Redis   Redis
	Sentry  Sentry
	Statsd  Statsd
}

type Http struct {
	Host string
	Port int64
}

type Casbin struct {
	AutoLoad         bool
	AutoLoadInternal int64
	Model            string
}

type MySQLConf struct {
	Host       string
	Port       int64
	DBName     string
	User       string
	Password   string
	Parameters string
}

type MySQL map[string]MySQLConf

type RedisConf struct {
	Host     string
	Port     int64
	User     string
	Password string
	DB       int64
	Timeout  int64
}

type Redis map[string]RedisConf

type Sentry struct {
	Dsn   string
	Level int64
}

type Statsd struct {
	Host   string
	Port   int
	Prefix string
	Name   string
}
