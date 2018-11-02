package db

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"codegen/config"
	"os"
)

type ForeignKey struct{
	TableName string
	ColName string
	RefTableName string
	RefColName string
}

var TableFKs map[string][]*ForeignKey


//初始化所有的外键
func initForeignKeys(){
	o := orm.NewOrm()
	sql := fmt.Sprintf("select table_name,column_name col_name,REFERENCED_TABLE_NAME ref_table_name,REFERENCED_COLUMN_NAME ref_col_name" +
		" from INFORMATION_SCHEMA.KEY_COLUMN_USAGE where table_name like '%s%%'" +
		" and  REFERENCED_TABLE_NAME is not null and REFERENCED_COLUMN_NAME is not null",
		config.Config.DB.TablePrefix)
	fmt.Println(sql)


	var foreignKeys []ForeignKey
	num,err := o.Raw(sql).QueryRows(&foreignKeys)
	if err != nil{
		println(err.Error())
		os.Exit(-1)
	}
	fmt.Printf("%d foreign keys in total!\n", num)
	fmt.Println(foreignKeys)

	TableFKs = make(map[string][]*ForeignKey)
	for _,item := range foreignKeys{
		TableFKs[item.TableName] = append(TableFKs[item.TableName], &item)
	}
}

