package run

import (
	"context"
	"log"

	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type DBModel interface {
	GetCasbin(ctx context.Context) ([]string, error)
}

// CasbinAdapter 重新定义casbin适配器
type CasbinAdapter struct {
	RoleRepo       DBModel
	UserRoleRepo   DBModel
	RoleObjectRepo DBModel
}

// loadRolePolicy 加载角色资源 (p,role::role_id,object::type:object)
func (c *CasbinAdapter) LoadRoleObject(ctx context.Context, m casbinModel.Model) error {
	result, err := c.RoleObjectRepo.GetCasbin(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, line := range result {
		persist.LoadPolicyLine(line, m)
	}
	return nil
}

// 加载角色-用户关系 (g,user::role_id,role::role_id)
func (c *CasbinAdapter) LoadUserRole(ctx context.Context, m casbinModel.Model) error {
	result, err := c.UserRoleRepo.GetCasbin(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, line := range result {
		persist.LoadPolicyLine(line, m)
	}

	return nil
}

// 加载角色之间的关系(g,role::parent_role_id,role::role_id)
func (c *CasbinAdapter) LoadRoleRelationship(ctx context.Context, m casbinModel.Model) error {
	result, err := c.RoleRepo.GetCasbin(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, line := range result {
		persist.LoadPolicyLine(line, m)
	}
	return nil
}

//  LoadPolicy从存储中加载所有策略规则（目前只做加载策略）
func (c *CasbinAdapter) LoadPolicy(m casbinModel.Model) error {
	ctx := context.Background()
	err := c.LoadRoleRelationship(ctx, m)
	if err != nil {
		log.Fatal("角色关系加载失败：", err)
	}
	err = c.LoadUserRole(ctx, m)
	if err != nil {
		log.Fatal("用户角色加载失败：", err)
	}
	err = c.LoadRoleObject(ctx, m)
	if err != nil {
		log.Fatal("角色资源加载失败：", err)
	}
	return nil
}

// SavePolicy 将所有策略规则保存到存储中
func (c *CasbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy 将策略规则添加到存储中。
// 这是自动保存功能的一部分。
func (c *CasbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy 从存储中删除政策规则。
// 这是自动保存功能的一部分。
func (c *CasbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy 删除从存储中匹配过滤器的策略规则。
// 这是自动保存功能的一部分。
func (c *CasbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
