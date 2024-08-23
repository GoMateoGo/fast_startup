package dbmysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var (
	DB *xorm.Engine
)

type MySql struct {
	Host   string
	User   string
	Pwd    string
	DbName string
	Show   bool
	Idle   int
	Open   int
	Config string
}

func mysqlDsn(m *MySql) string {
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Pwd, m.Host, m.DbName)
}

// 创建db句柄公共函数
func NewMysql(m *MySql) {
	dsn := mysqlDsn(m)
	fmt.Println("链接信息:", dsn)
	dbHandle, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("创建db引擎错误", err)
		panic(err.Error())
	}
	if err := dbHandle.Ping(); err != nil {
		fmt.Println("测试连接错误!", err)
		panic(err.Error())
	}
	dbHandle.SetTableMapper(core.SameMapper{})
	dbHandle.SetColumnMapper(core.SnakeMapper{})
	dbHandle.ShowSQL(m.Show)
	dbHandle.SetMaxIdleConns(m.Idle)
	dbHandle.SetMaxOpenConns(m.Open)
	DB = dbHandle
}
