package role

import (
	"context"
	"fmt"
	"role-system/settings"
)

// GetRoles
func (r Role) GetCasbin(ctx context.Context) ([]string, error) {
	var dbConn = settings.GetMysqlServer("default")
	var roles []Role
	err := dbConn.SelectContext(ctx, &roles, "select * from role")
	if err != nil {
		return nil, err
	}
	var result []string
	for _, role := range roles {
		line := fmt.Sprintf(settings.RoleRelationshipPolicy, role.Pid, role.ID)
		result = append(result, line)
	}
	return result, nil
}
