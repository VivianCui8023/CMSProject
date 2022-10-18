package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"webProject/model"
	"xorm.io/core"
)

func Engine_sql() *xorm.Engine {
	//第一个参数是驱动名字，第二个参数是：用户名：密码@/数据库名字?charset=utf8
	engine, err := xorm.NewEngine("mysql", "root:root@/cmsdata?charset=utf8")

	//defer engine.Close()
	//设置对象 与数据库映射，支持自动建表
	err = engine.CreateTables(new(model.Admin))
	engine.SetMapper(core.GonicMapper{})
	err = engine.Sync2(
		new(model.Admin),
		new(model.User),
		new(model.UserOrder))
	if err != nil {
		panic(err.Error())
	}
	//打印查询语句
	engine.ShowSQL(true)
	//设置最大连接数
	engine.SetMaxIdleConns(10)
	return engine

}
