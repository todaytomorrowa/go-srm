package util

import (
    "io/ioutil"
    "fmt"
    "encoding/json"
    "flag"
)

var (
    instance *Config
    conf = flag.String("conf","../etc/config.json","描述")

)

type Config struct {
    BasePath string `json:"base_path"`
    DataPath string `json:"data_path"`
}

func init() {
    //NewConfigWithFile("/Users/derek/go/src/sixedu/data/config.json")
    flag.Parse()
    NewConfigWithFile(*conf)
}

func NewConfigWithFile(name string) {

    if instance == nil {
        c  := &Config{}
        instance = c


    fielContentsBytes, err := ioutil.ReadFile(name)
    if err != nil {
        fmt.Println("读取配置文件失败：",err)
        return
    }

    err = json.Unmarshal(fielContentsBytes, c)

    if err != nil {
        fmt.Println("解析失败")
        return
    }

    }

    
  //  fmt.Println(fileContent)
}

func GetConfig() *Config{
    //fmt.Println(instance)
    //fmt.Printf(*conf)
    return instance
}