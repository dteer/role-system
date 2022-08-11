package run

import (
	"log"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

// 拼接角色命名(了解到的尚未满足现需求，暂时屏蔽)
func SplicingRole(args ...interface{}) (interface{}, error) {
	// key1 := args[0].(string)
	// key2 := args[1].(string)
	return true, nil
}

var casbinObj *casbin.SyncedEnforcer
var casbinClean func()

type CasbinConfig struct {
	Model string
}

// casbin初始化
func InitCasbin(adapter persist.Adapter, config CasbinConfig) (*casbin.SyncedEnforcer, func(), error) {
	// 初始化casbin适配器
	e, err := casbin.NewSyncedEnforcer(config.Model)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	// e.EnableLog(true)
	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	e.AddFunction("splicing_role", SplicingRole)
	e.EnableEnforce(true)
	// 自动加载casbin策略
	e.StartAutoLoadPolicy(time.Duration(3) * time.Second)
	casbinClean = func() {
		e.StopAutoLoadPolicy()
	}
	return e, casbinClean, nil
}

func GetCasbinObj() *casbin.SyncedEnforcer {
	return casbinObj
}
func GetCasbinClean() func() {
	return casbinClean
}
