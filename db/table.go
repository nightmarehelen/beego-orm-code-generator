package db

import (
	"github.com/astaxie/beego/orm"
	"codegen/config"
	"fmt"
	"os"
)

type Table struct{
	Name string `orm:column("name")`
	Comment string `orm:column("comment")`
	Columns map[string]*Column
}

var Tables map[string]*Table

func initTables(){
	o :=orm.NewOrm()
	o.Using(config.Config.DB.Alias)

	sql := fmt.Sprintf("select table_name name,table_comment comment from information_schema.`TABLES` t " +
		"where t.table_name like '%s%%'", config.Config.DB.TablePrefix)
	fmt.Println(sql)
	var tables []Table
	num,err := o.Raw(sql).QueryRows(&tables)
	if err != nil{
		println(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("%d tables in total!\n", num)
	fmt.Println(tables)

	Tables = make(map[string]*Table)
	for _,item := range tables{
		Tables[item.Name] = &item
	}

	initColumns()
}