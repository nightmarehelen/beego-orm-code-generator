package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"fmt"
)

var Config *config
type config struct{
	DB *dbConfig `yaml:DB`

	ModelTemplatePath string `yaml:"model_template_path"`
	ModelOutputPath string `yaml:"model_output_path"`
	ServiceTemplatePath string `yaml:"service_template_path"`
	ServiceOutputPath string `yaml:"service_output_path"`
	Tables []string `yaml:"tables"`
}

type dbConfig struct{
	Alias string `yaml:"alias"`
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Encoding string `yaml:"encoding"`
	TablePrefix string `yaml:"table_prefix"`
}

func init(){
	bytes,err := ioutil.ReadFile("conf/app.yaml")
	if err != nil{
		panic(fmt.Errorf("fail to read config file : conf/app.yaml,caused by %s",err.Error()))
	}

	Config = &config{}
	err = yaml.Unmarshal(bytes, Config)
	if err != nil{
		panic(fmt.Errorf("fail to parse : conf/app.yaml,caused by %s", err.Error()))
	}

	bytes,err = json.Marshal(Config)
	fmt.Println(string(bytes))
}
