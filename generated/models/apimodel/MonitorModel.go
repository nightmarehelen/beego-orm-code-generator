package apimodel

import (
	"github.com/astaxie/beego/validation"
	"regexp"
	"github.com/astaxie/beego/orm"

)

func init(){
	//注册模型至系统
	orm.RegisterModel(new(Monitor))

	//正则编译
	//regName
    regName = regexp.MustCompile(`[0-9a-zA-Z]{4,6}`)
    
}

//校验正则声明
//regName
var regName *regexp.Regexp

//Monitor 监控设备信息
type Monitor struct{ 
    //Description  
    Description string `json:"Description" orm:"column(description);"`
    //Id  
    Id int `json:"Id" orm:"column(id);"`
    //Name  {    "comment":"我是注释",    "regex":"[0-9a-zA-Z]{4,6}",    "error_msg":"请输入4-6位大小写字母及数字"}
    Name string `json:"Name" orm:"column(name);"`
    //Resource  
    Resource string `json:"Resource" orm:"column(resource);"`
    //Type  
    Type int `json:"Type" orm:"column(type);"`
    //X  
    X float64 `json:"X" orm:"column(x);"`
    //Y  
    Y float64 `json:"Y" orm:"column(y);"`
    //Z  
    Z float64 `json:"Z" orm:"column(z);"`
	BaseModel
}

/**
系统根据该函数映射实体和数据库表的关系
 */
func (m *Monitor) TableName() string {
	return "ors_cg_monitor"
}

func (m *Monitor) IsValid() (bool, string) {
    //正则校验
	valid := validation.Validation{}
    //regName
    var regName *regexp.Regexp
    valid.Match(m.Name, regName, "Name").Message("请输入4-6位大小写字母及数字")
    
	return m.SprintfValidation(&valid)
}