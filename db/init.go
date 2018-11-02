package db

import (
	"codegen/config"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//元数据库，查询系统当中的表信息
const InformationSchema = "INFORMATION_SCHEMA"

func init(){
	initDB()
	initTables()
	initForeignKeys()
}


func initDB(){

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Config.DB.Username,
		config.Config.DB.Password,
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.Database,
		config.Config.DB.Encoding)

	fmt.Println(connStr)

	orm.RegisterDataBase(config.Config.DB.Alias, config.Config.DB.Type, connStr, 30)
}
