package apimodel

import (
	"github.com/astaxie/beego/validation"
	"strings"
	"codegen/config"
)

type BaseModel struct{

}

//SprintfValidation 格式化校验信息
func (*BaseModel) SprintfValidation(valid *validation.Validation) (bool,string){
	if !valid.HasErrors(){
		return true, ""
	}
	sb := strings.Builder{}
	for _, err := range valid.Errors {
		sb.WriteString(err.Key)
		sb.WriteString(":")
		sb.WriteString(err.Message)
		sb.WriteString("\n")
	}
	return false,sb.String()
}


//TableNameImpl 系统根据该函数映射实体和数据库表的关系
func (* BaseModel) TableNameImpl(model string) string{
	prefix := config.Config.DB.TablePrefix
	return prefix + model
}
