package settings

import (
	configs "role-system/configs"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

// 常量配置
var Conf = &configs.Conf

// 日志配置
var Logger zerolog.Logger

// mysql服务
var MysqlServer = make(map[string]*sqlx.DB)

// redis服务
var RedisDB = make(map[string]*redigo.Pool)

// gin 服务
var Engine *gin.Engine

// casbin服务
var Casbin *casbin.SyncedEnforcer

// 定义策略规则
var RoleObjectPolicy = "p,role::%d,object::%d:%s,%s" // 角色id，资源类型，资源，访问动作
var UserRolePolicy = "g,user::%d,role::%d"           // 用户id，角色id
var RoleRelationshipPolicy = "g,role::%d,role::%d"   // 角色id，角色id
