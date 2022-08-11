package user

type UserRole struct {
	ID         int64  `db:"id"`
	RoleId     int64  `db:"role_id"`
	UserId     int64  `db:"user_id"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
