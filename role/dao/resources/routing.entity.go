package resources

type Routing struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Pid        int64  `db:"pid"`
	Tid        int64  `db:"tid"`
	Path       string `db:"path"`
	Desc       string `db:"desc"`
	Type       int64  `db:"type"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}
