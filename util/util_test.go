package util

import (
    "testing"
    "fmt"
    "flag"
)

func TestConfig(t *testing.T) {
  //  NewConfigWithFile("/Users/derek/go/src/sixedu/data/config.json")
    //  c := GetConfig()
      fmt.Println("")

    args := []string{
        "-conf=这是测试命令行参数",
    }
    flag.CommandLine.Parse(args)
    GetConfig()
    //  fmt.Println(c)
    
 }