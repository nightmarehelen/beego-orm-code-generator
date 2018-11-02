package util

import (
	"encoding/json"
	"bytes"
)

func IndentJSON(v interface{}) string{

	data, err := json.Marshal(v)
	if err != nil{
		panic(err)
	}
	buffer  := bytes.NewBuffer(make([]byte,0,0))

	err = json.Indent(buffer, data, "", "\t")
	if err != nil{
		panic(err)
	}
	return buffer.String()
}

//ToCamelBak 将下划线分隔的标识符转化为驼峰格式
func ToCamelBak(input string) string{
	//空字符串返回
	if input == ""{
		return input
	}
	//拷贝字符串
	buffer := make([]byte, 0, 0)
	origin := []byte(input)

	//下划线后面的字母大写
	for idx,b := range origin{
		if b == '_' {
			continue
		}
		//下划线后面有内容
		if (idx-1 >=0 && origin[idx-1] == '_') || idx == 0{
			buffer = append(buffer, toUpper(origin[idx]))
			continue
		}
		//直接追加
		buffer = append(buffer, origin[idx])
	}
	return string(buffer)
}

//ToCamelBak 将下划线分隔的标识符转化为驼峰格式,首字母小写
func ToCamelBak2(input string) string{
	//空字符串返回
	if input == ""{
		return input
	}
	//拷贝字符串
	buffer := make([]byte, 0, 0)
	origin := []byte(input)

	//下划线后面的字母大写
	for idx,b := range origin{
		if b == '_' {
			continue
		}

		if idx == 0 {
			buffer = append(buffer, toLower(origin[idx]))
		}
		//下划线后面有内容
		if idx-1 >=0 && origin[idx-1] == '_'{
			buffer = append(buffer, toUpper(origin[idx]))
			continue
		}
		//直接追加
		buffer = append(buffer, origin[idx])
	}
	return string(buffer)
}

func toUpper(b byte) byte{
	//首字母大写
	if b >= 'a' && b <= 'z'{
		b = b + 'A' - 'a'
	}
	return b
}

func toLower(b byte) byte{
	//首字母大写
	if b >= 'A' && b <= 'Z'{
		b = b - 'A' - 'a'
	}
	return b
}

