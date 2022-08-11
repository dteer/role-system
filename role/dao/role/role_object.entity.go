package role

type RoleObject struct {
	ID         int64  `db:"id"`
	ObjectType int64  `db:"object_type"`
	ObjectId   int64  `db:"object_id"`
	Action     string `db:"action"`
	RoleId     int64  `db:"role_id"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
