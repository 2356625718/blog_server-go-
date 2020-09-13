package tool

import(
	"bufio"
	"encoding/json"
	"log"
	"os"
)
//配置信息结构体
type Config struct{
	AppName string `json:"appName"`
	AppMode string `json:"appMode"`
	AppHost string `json:"appHost"`
	AppPort string `json:"appPort"`
	Db `json:"database"`
}
//数据库结构体
type Db struct{
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	Db_name string `json:"db_name"`
	Charset string `json:"charset"`
	Show_sql bool `json:"show_sql"`
}
//配置结构体工厂函数
func ParseConfig(path string)(*Config,error){
	var cfg *Config
	f,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&cfg)
	if err != nil{
		log.Fatal(err)
		return nil,err
	}
	return cfg,nil
}
