package apisrv

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"reflect"
	"time"
)

type BaseService struct {
	//Orm对象，用于控制事务处理
	O orm.Ormer
}

//反射获取接口的所有字段，如果字段值不为默认值，则认为该字段需要更新
func (s *BaseService) GetColumns2Update(obj interface{}) ([]string){
	//反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	t := reflect.TypeOf(obj)
	//调用t.Name方法来获取这个类型的名称
	fmt.Println("Type:", t.Name())

	//打印出所包含的字段
	v := reflect.ValueOf(obj)
	fmt.Println("Fields:")

	columns := make([]string, 0, 0)
	//通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
	for i := 0; i < t.NumField(); i++ {
		//通过这个i作为它的索引，从0开始来取得它的字段
		f := t.Field(i)
		//通过interface方法来取出这个字段所对应的值
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v =%v\n", f.Name, f.Type, val)
		//目前只处理数字，字符串，时间三种格式，未完待续
		switch v.Kind() {
			case reflect.String:
				if val != ""{
					columns = append(columns, f.Name)
				}
			case reflect.Float32,reflect.Float64,reflect.Int,reflect.Int64:
				if val != 0{
					columns = append(columns, f.Name)
				}
			case reflect.Ptr:
				if val != nil {
					_,ok := v.Field(i).Interface().(time.Time)
					if ok{
						columns = append(columns, f.Name)
					}
				}
		}
	}
	return columns
}
