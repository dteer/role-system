package user

type User struct {
	ID         int64  `db:"column:id"`
	Name       string `db:"column:name"`
	CreateTime string `db:"column:create_time"`
	UpdateTime string `db:"column:update_time"`
}
