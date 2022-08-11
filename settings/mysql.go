package settings

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

func MysqlConnect(name string) (*sqlx.DB, error) {
	mysqlCof := Conf.MySQL[name]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlCof.User, mysqlCof.Password, mysqlCof.Host, mysqlCof.Port, mysqlCof.DBName)
	if mysqlCof.Parameters != "" {
		dsn = fmt.Sprintf("%s?%s", dsn, mysqlCof.Parameters)
	}
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 设置连接的最大生存时间，以确保连接可以被驱动安全关闭。官方建议小于5分钟。
	db.SetConnMaxLifetime(time.Minute * 3)
	// 设置打开的最大连接数，取决于mysql服务器和具体应用程序
	db.SetMaxOpenConns(1000)
	// 设置最大闲置连接数，这个连接数应大于等于打开的最大连接，否则需要额外连接时会频繁进行打开关闭。
	// 最好与最大连接数保持相同，当大于最大连接数时，内部自动会减少到与最大连接数相同。
	db.SetMaxIdleConns(1000)

	// 设置闲置连接的最大存在时间, support>=go1.15
	db.SetConnMaxIdleTime(time.Minute * 3)
	return db, nil
}

func InitMysqlServer() {
	for dbName := range Conf.MySQL {
		db, err := MysqlConnect(dbName)
		if err != nil {
			log.Fatal(err)
			continue
		}
		MysqlServer[dbName] = db
	}
}

func GetMysqlServer(dbName string) *sqlx.DB {
	var mu sync.RWMutex
	mu.Lock()
	server, ok := MysqlServer[dbName]
	if !ok {
		var err error
		server, err = MysqlConnect(dbName)
		if err != nil {
			panic(fmt.Sprintf("数据库连接异常：%s", err.Error()))
		}
		MysqlServer[dbName] = server
	}
	mu.Unlock()
	return server
}
