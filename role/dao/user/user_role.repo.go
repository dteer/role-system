package user

import (
	"context"
	"fmt"
	"role-system/settings"
)

// 获取用户和角色关系
func (UserRole) GetCasbin(ctx context.Context) ([]string, error) {
	var dbConn = settings.GetMysqlServer("default")
	var roles []UserRole
	err := dbConn.SelectContext(ctx, &roles, "select * from user_role")
	if err != nil {
		return nil, err
	}
	var result []string
	for _, obj := range roles {
		line := fmt.Sprintf(settings.UserRolePolicy, obj.UserId, obj.RoleId)
		result = append(result, line)
	}
	return result, nil
}
