package apimodel

import (
	"time"
	"github.com/astaxie/beego/validation"
	"regexp"
	"github.com/astaxie/beego/orm"

)

func init(){
	//注册模型至系统
	orm.RegisterModel(new(LevelItem))

	//正则编译
	//regCode
    regCode = regexp.MustCompile(`[0-9a-zA-Z]{4,6}`)
    //regText
    regText = regexp.MustCompile(`[0-9a-zA-Z]{4,6}`)
    
}

//校验正则声明
//regCode
var regCode *regexp.Regexp//regText
var regText *regexp.Regexp

//LevelItem 楼层附属物
type LevelItem struct{ 
    //Blob  
    Blob []byte `json:"Blob" orm:"column(blob);type(blob);"`
    //Code  {    "comment":"我是注释",    "regex":"[0-9a-zA-Z]{4,6}",    "error_msg":"请输入4-6位大小写字母及数字"}
    Code string `json:"Code" orm:"column(code);"`
    //Date  
    Date time.Time `json:"Date" orm:"column(date);"`
    //Datetime  
    Datetime *time.Time `json:"Datetime" orm:"column(datetime);"`
    //Enum  
    Enum string `json:"Enum" orm:"column(enum);"`
    //Id  
    Id int `json:"Id" orm:"column(id);"`
    //LevelId  
    LevelId int `json:"LevelId" orm:"column(level_id);"`
    //RateOfProgress  
    RateOfProgress float32 `json:"RateOfProgress" orm:"column(rate_of_progress);"`
    //Text  {    "comment":"我是注释",    "regex":"[0-9a-zA-Z]{4,6}",    "error_msg":"请输入4-6位大小写字母及数字"}
    Text string `json:"Text" orm:"column(text);"`
    //Timestamp  
    Timestamp time.Time `json:"Timestamp" orm:"column(timestamp);"`
    //Type  
    Type string `json:"Type" orm:"column(type);"`
	BaseModel
}

/**
系统根据该函数映射实体和数据库表的关系
 */
func (m *LevelItem) TableName() string {
	return "ors_cg_level_item"
}

func (m *LevelItem) IsValid() (bool, string) {
    //正则校验
	valid := validation.Validation{}
    //regCode
    var regCode *regexp.Regexp
    valid.Match(m.Code, regCode, "Code").Message("请输入4-6位大小写字母及数字")
    //regText
    var regText *regexp.Regexp
    valid.Match(m.Text, regText, "Text").Message("请输入4-6位大小写字母及数字")
    
	return m.SprintfValidation(&valid)
}