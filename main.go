package main

import (
	"fmt"
	configs "role-system/configs"
	"role-system/role/dao/role"
	"role-system/role/dao/user"
	"role-system/settings"
	runCasbin "role-system/settings/casbin"
)

func main() {
	// 配置初始化``
	configs.InitConfig()
	// 日志初始化
	settings.InitLog()
	// 数据库初始化
	settings.InitMysqlServer()
	// redis初始化
	settings.InitRedis()
	// casbin初始化
	dapter := runCasbin.CasbinAdapter{
		RoleRepo:       role.Role{},
		UserRoleRepo:   user.UserRole{},
		RoleObjectRepo: role.RoleObject{},
	}
	cabinConfig := runCasbin.CasbinConfig{
		Model: settings.Conf.Casbin.Model,
	}
	settings.Casbin, _, _ = runCasbin.InitCasbin(&dapter, cabinConfig)
	// gin项目启动
	router := settings.InitRouter()
	addr := fmt.Sprintf("%s:%d", configs.Conf.HTTP.Host, configs.Conf.HTTP.Port)
	router.Run(addr)

}
