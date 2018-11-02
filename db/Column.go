package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"os"
	"codegen/util"
	"encoding/json"
)

//类型转换字典
var TypeDict map[string]string

var NumTypeDict map[string]string

var StrTypeDict map[string]string
//全局变量，存储所有列
type columns map[string][]*Column


type Column struct{
	TableName string
	ColumnName string `orm:"column(column_name)"`
	ColumnDefault interface{}
	IsNull string `orm:"column(is_null)"`
	DataType string
	CharacterMaximumLength int
	NumericPrecision int
	NumericScale int
	ColumnType string
	//原始的数据库注释
	ColumnComment string
	//格式化解析后的注释，包含的正则校验等内容
	Comment *Comment
}

func (c *Column)parseComment(){
	if c.ColumnComment == ""{
		return
	}
	cmt := Comment{}
	err := json.Unmarshal([]byte(c.ColumnComment), &cmt)
	if err != nil{
		println(fmt.Sprintf("comment of column %s of table %s is not well-formatted json:%s,caused by %s",
			c.ColumnName, c.TableName, c.ColumnComment, err))
		return
	}
	c.Comment = &cmt
}
func initTypeDict(){
	TypeDict = make(map[string]string)

	//mysql数据类型到go的类型映射
	TypeDict["varchar"] = "string"
	TypeDict["date"] = "time.Time"
	TypeDict["float"] = "float"
	TypeDict["time"] = "time.Time"
	TypeDict["mediumblob"] = "[]byte"
	TypeDict["smallint"] = "int"
	TypeDict["longblob"] = "[]byte"
	TypeDict["double"] = "float64"
	TypeDict["decimal"] = "float64"
	TypeDict["blob"] = "[]byte"
	TypeDict["varbinary"] = "[]byte"
	TypeDict["binary"] = "[]byte"
	TypeDict["char"] = "string"
	TypeDict["set"] = "map([]interface,bool)"
	TypeDict["datetime"] = "time.Time"
	TypeDict["json"] = "string"
	TypeDict["mediumtext"] = "string"
	TypeDict["longtext"] = "string"
	TypeDict["text"] = "string"
	TypeDict["enum"] = "string"
	TypeDict["int"] = "int"
	TypeDict["bigint"] = "int64"
	TypeDict["timestamp"] = "time.Time"
	TypeDict["varchar"] = "string"

	//数字类型
	NumTypeDict = make(map[string]string)
	NumTypeDict["float"] = "float"
	NumTypeDict["smallint"] = "int"
	NumTypeDict["longblob"] = "[]byte"
	NumTypeDict["double"] = "float64"
	NumTypeDict["decimal"] = "float64"
	NumTypeDict["int"] = "int"
	NumTypeDict["bigint"] = "int64"

	//字符类型，需要设置长度的，可能需要校验长度
	StrTypeDict = make(map[string]string)
	StrTypeDict["varchar"] = "string"
	StrTypeDict["mediumblob"] = "[]byte"
	StrTypeDict["longblob"] = "[]byte"
	StrTypeDict["blob"] = "[]byte"
	StrTypeDict["varbinary"] = "[]byte"
	StrTypeDict["binary"] = "[]byte"
	StrTypeDict["char"] = "string"
	StrTypeDict["datetime"] = "time.Time"
	StrTypeDict["json"] = "string"
	StrTypeDict["mediumtext"] = "string"
	StrTypeDict["longtext"] = "string"
	StrTypeDict["text"] = "string"
	StrTypeDict["varchar"] = "string"
}

func initColumns(){
	initTypeDict()

	o := orm.NewOrm()
	for name,table := range Tables{

		table.Columns = make(map[string]*Column)

		sql := fmt.Sprintf("select table_name `table_name`,column_name `column_name`,column_default `column_default`," +
			" is_nullable `is_null`,data_type `data_type`, character_maximum_length `character_maximum_length`," +
			" numeric_precision `numeric_precision`,numeric_scale `numeric_scale`, " +
			" column_type `column_type`,column_comment `column_comment`" +
			" from information_schema.`COLUMNS` t where t.table_name='%s'", name)
		fmt.Println(sql)
		var columns []Column
		num,err := o.Raw(sql).QueryRows(&columns)
		if err != nil{
			println(err.Error())
			os.Exit(-1)
		}
		fmt.Printf("load %d columns for table %s\n", num, name)

		for _,column := range columns{
			column.parseComment()
			table.Columns[column.ColumnName] = &column
		}

		fmt.Println(util.IndentJSON(table))
	}
}