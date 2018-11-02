package db

import (
	"github.com/astaxie/beego/orm"
	"codegen/config"
	"fmt"
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
	tables := make([]*Table, 0, 0)
	num,err := o.Raw(sql).QueryRows(&tables)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%d tables in total!\n", num)
	fmt.Println(tables)

	Tables = make(map[string]*Table)
	for _,item := range tables{
		if needHandle(item.Name){
			Tables[item.Name] = item
		}

	}

	initColumns()
}

func needHandle(tableName string) bool{
	for _,item := range config.Config.Tables{
		if item == tableName{
			return true
		}
	}
	return false
}