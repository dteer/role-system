package role

import (
	"context"
	"fmt"
	"log"
	"role-system/settings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type RoleObjectCasbin struct {
	RoleId     int64  `db:"role_id"`
	Action     string `db:"action"`
	ObjectId   int64  `db:"object_id"`
	Object     string `db:"object"`
	ObjectType int8   `db:"object_type"`
}

// 获取用户和角色关系
func (RoleObject) GetCasbin(ctx context.Context) ([]string, error) {
	var dbConn = settings.GetMysqlServer("default")
	sql_routing := `
			SELECT  role_object.role_id AS role_id, role_object.action AS action,
					role_object.object_type AS object_type,
					routing.path AS object ,routing.id AS object_id
			FROM role_object
			INNER JOIN routing ON role_object.object_id = routing.id
			where role_object.object_type = 1
	`
	sqls := []string{sql_routing}
	// TODO 异步处理多请求，未经过验证是否安全
	var channel = make(chan []RoleObjectCasbin, len(sqls))
	var wg sync.WaitGroup // 定义一个同步等待的组
	for _, sql := range sqls {
		wg.Add(1)
		sql := sql
		go func(ctx context.Context, sql string, channel chan []RoleObjectCasbin) {
			var roles []RoleObjectCasbin
			err := dbConn.SelectContext(ctx, &roles, sql)
			if err != nil {
				log.Fatal(err)
				wg.Done()
				return
			}
			channel <- roles
			wg.Done()
		}(ctx, sql, channel)
	}
	wg.Wait()
	close(channel)
	var objects []RoleObjectCasbin
	for object := range channel {
		objects = append(objects, object...)
	}
	var results []string
	for _, obj := range objects {
		line := fmt.Sprintf(settings.RoleObjectPolicy, obj.RoleId, obj.ObjectType, obj.Object, obj.Action)
		results = append(results, line)
	}
	return results, nil
}
