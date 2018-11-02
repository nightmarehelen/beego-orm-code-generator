package api

import (
	"time"
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
var regCode *regexp.Regexp
//regText
var regText *regexp.Regexp


//LevelItem 楼层附属物
type LevelItem struct{
    
    //Blob  
    Blob []byte `json:"Blob" orm:"column(blob);null;"`
    //Code  {    "comment":"我是注释",    "regex":"[0-9a-zA-Z]{4,6}",    "error_msg":"请输入4-6位大小写字母及数字"}
    Code string `json:"Code" orm:"column(code);null;"`
    //Date  
    Date time.Time `json:"Date" orm:"column(date);null;"`
    //Datetime  
    Datetime time.Time `json:"Datetime" orm:"column(datetime);null;"`
    //Enum  
    Enum string `json:"Enum" orm:"column(enum);null;"`
    //Id  
    Id int `json:"Id" orm:"column(id);"`
    //LevelId  
    LevelId int `json:"LevelId" orm:"column(level_id);null;"`
    //RateOfProgress  
    RateOfProgress float32 `json:"RateOfProgress" orm:"column(rate_of_progress);null;"`
    //Text  {    "comment":"我是注释",    "regex":"[0-9a-zA-Z]{4,6}",    "error_msg":"请输入4-6位大小写字母及数字"}
    Text string `json:"Text" orm:"column(text);null;"`
    //Timestamp  
    Timestamp time.Time `json:"Timestamp" orm:"column(timestamp);null;"`
    //Type  
    Type string `json:"Type" orm:"column(type);null;"`
	BaseModel
}

/**
系统根据该函数映射实体和数据库表的关系
 */
func (m *LevelItem) TableName() string {
	return "ors_cg_level_item"
}


func (m *LevelItem) IsValid() (bool, string) {

	/**
	valid := validation.Validation{}

	valid.Match(user.Username, regUsername, "username").Message("请输入正确的用户名,4到16位（字母，数字，下划线，减号）")

	valid.Match(user.Password, regPassword, "password").Message("请输入正确的用户名,4到16位（字母，数字，下划线，减号）")

	valid.Email(user.Email, "email").Message("邮箱地址有误")

	return user.SprintfValidation(&valid)*/
	return false, ""
}