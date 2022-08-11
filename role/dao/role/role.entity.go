package role

type Role struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Pid        int64  `db:"pid"`
	Tid        int64  `db:"tid"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
