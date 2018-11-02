package service

import "github.com/astaxie/beego/orm"

type BaseService struct{
	//Orm对象，用于控制事务处理
	O orm.Ormer
}



