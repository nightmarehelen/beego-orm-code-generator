package db

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"codegen/util"
	"encoding/json"
	"strings"
)

//类型转换字典
var TypeDict map[string]string

var NumTypeDict map[string]string

var StrTypeDict map[string]string

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
	ColumnComment string `orm:"column(column_comment);type(text)"`
	ColumnKey string `orm:"column(column_key)"`
	//格式化解析后的注释，包含的正则校验等内容
	Comment *Comment
}


type Comment struct{
	Comment string `json:"comment"`
	Regex string `json:"regex"`
	ErrorMsg string `json:"error_msg"`
}


func (c *Column)parseComment(){
	if c.ColumnComment == ""{
		return
	}

	c.ColumnComment = strings.Replace(c.ColumnComment, "\r", "", -1)
	c.ColumnComment = strings.Replace(c.ColumnComment, "\n", "", -1)
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
	TypeDict["float"] = "float32"
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
	TypeDict["datetime"] = "*time.Time"
	TypeDict["json"] = "string"
	TypeDict["mediumtext"] = "string"
	TypeDict["longtext"] = "string"
	TypeDict["text"] = "string"
	TypeDict["enum"] = "string"
	TypeDict["int"] = "int"
	TypeDict["bigint"] = "int64"
	TypeDict["tinyint"] = "int"
	TypeDict["timestamp"] = "*time.Time"
	TypeDict["varchar"] = "string"

	//数字类型
	NumTypeDict = make(map[string]string)
	NumTypeDict["float"] = "float32"
	NumTypeDict["smallint"] = "int"
	NumTypeDict["longblob"] = "[]byte"
	NumTypeDict["double"] = "float64"
	NumTypeDict["decimal"] = "float64"
	NumTypeDict["int"] = "int"
	NumTypeDict["bigint"] = "int64"
	NumTypeDict["tinyint"] = "int"

	//字符类型，需要设置长度的，可能需要校验长度
	StrTypeDict = make(map[string]string)
	StrTypeDict["varchar"] = "string"
	StrTypeDict["mediumblob"] = "[]byte"
	StrTypeDict["longblob"] = "[]byte"
	StrTypeDict["blob"] = "[]byte"
	StrTypeDict["varbinary"] = "[]byte"
	StrTypeDict["binary"] = "[]byte"
	StrTypeDict["char"] = "string"
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
			" column_type `column_type`,column_comment `column_comment`,column_key `column_key` " +
			" from information_schema.`COLUMNS` t where t.table_name='%s'", name)
		fmt.Println(sql)
		columns := make([]*Column, 0, 0)
		num,err := o.Raw(sql).QueryRows(&columns)
		if err != nil{
			panic(err.Error())
		}

		fmt.Printf("load %d columns for table %s\n", num, name)

		for _,column := range columns{
			column.parseComment()
			column.TableName = name
			table.Columns[column.ColumnName] = column
		}

		fmt.Println(util.IndentJSON(table))
	}
}


func MySQLType2Go(mysql string) string{
	t, ok := TypeDict[mysql]
	if !ok{

		panic(fmt.Errorf("unsurpporte type of mysql to map to go :%s", mysql))
	}
	return  t
}