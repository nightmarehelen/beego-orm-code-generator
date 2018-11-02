package api

import (
	"time"
	"github.com/astaxie/beego/validation"
	"regexp"
	"github.com/astaxie/beego/orm"

)


func init(){

	//注册模型至系统
	orm.RegisterModel(new(Monitor))

	//正则编译
	regUsername = regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)

	regPassword = regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
}

var regUsername  *regexp.Regexp
var regPassword  *regexp.Regexp
var regEmail  *regexp.Regexp


//Monitor 树木信息
type Monitor struct{
	Id int  `json:"Id,omitempty" orm:"auto;size(11);column(Id);pk;"`
	Username string `json:"Username,omitempty" orm:"size(20);column(username);unique;"`
	Password string `json:"Password,omitempty" orm:"size(32);column(password);unique;"`
	Email string `json:"Email,omitempty" orm:"size(32);column(email);unique;"`
	LastLogin time.Time `orm:"column(last_login);null;type(datetime)"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Updated time.Time `orm:"column(updated);null;type(datetime);auto_now"`

    
        id
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        local_eular_angles_y
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        local_scale
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        name
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        season
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        type
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        x
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        y
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    
        z
        {ors_cg_tree z &lt;nil&gt; YES double 0 17 6 double(17,6)  &lt;nil&gt;}

    

	BaseModel
}

/**
系统根据该函数映射实体和数据库表的关系
 */
func (user *User) TableName() string {
	return user.TableNameImpl("user")
}


func (user *User) IsValid() (bool, string) {

	valid := validation.Validation{}

	valid.Match(user.Username, regUsername, "username").Message("请输入正确的用户名,4到16位（字母，数字，下划线，减号）")

	valid.Match(user.Password, regPassword, "password").Message("请输入正确的用户名,4到16位（字母，数字，下划线，减号）")

	valid.Email(user.Email, "email").Message("邮箱地址有误")

	return user.SprintfValidation(&valid)
}