package gen

import (
	"codegen/db"
	"fmt"
	"strings"
	"codegen/config"
	"codegen/util"
	"io/ioutil"
	"text/template"
	"os"
)

//生成模型文件
func GenerateModels(){
	//获取解析模板
	t := getTemplateParser(config.Config.ModelTemplatePath)

	for tableName,table := range db.Tables {

		found := false
		for _,temp := range config.Config.Tables{
			if temp == tableName{
				found = true
			}
		}
		if !found {
			continue
		}

		//去掉表前缀，改成驼峰形式，并且文件名加上Model
		modelName := util.ToCamelBak(strings.Replace(tableName, config.Config.DB.TablePrefix, "", -1))
		outputFileName := modelName + "Model"

		outputFileName = pathJoin(config.Config.ModelOutputPath, outputFileName) +".go"

		file, err := os.Create(outputFileName)
		if err != nil {
			panic(fmt.Errorf("fail to create file %s,caused by %s", outputFileName, err))
		}

		model := &Model{}

		model.PackageName = getPackageName(config.Config.ModelOutputPath)
		model.ModelName = modelName
		model.Table = table

		fmt.Printf("generate model for table %s, contents: %s\n", model.Table.Name, util.IndentJSON(model.Table))
		t.Execute(file, model)
	}
}


func pathJoin(path string, fileName string) string{
	//如果以斜杠结尾
	if strings.LastIndex(path,"/") == len(path) -1{
		fileName = path + fileName
	}else{
		fileName = path + "/" + fileName
	}
	return fileName
}

func getPackageName(path string) string{
	//如果以斜杠结尾
	if strings.LastIndex(path,"/") == len(path) -1{
		idx := strings.LastIndex(path[0:len(path)-1], "/")
		if idx == -1 {
			panic(fmt.Errorf("invalid path to generate package name :%s", path))
		}
		return path[idx+1:len(path)-1]
	}else{
		idx := strings.LastIndex(path, "/")
		if idx == -1 {
			panic(fmt.Errorf("invalid path to generate package name :%s", path))
		}
		return path[idx+1:]
	}
}

func getFuncMap() map[string]interface{}{
	funcMap := make(map[string]interface{})
	//驼峰格式转换
	funcMap["ToCamelBak"] = util.ToCamelBak

	funcMap["MySQLType2Go"] = db.MySQLType2Go

	return funcMap
}


func getTemplateParser(templatePath string) * template.Template {
	//加载模板
	data, err := ioutil.ReadFile(templatePath)
	if err != nil{
		panic(fmt.Errorf("fail to load model templates:%s,caused by %s", templatePath, err))
	}

	t := template.Must(template.New("Model").Funcs(getFuncMap()).Parse(string(data)))
	return t
}
