package util

import (
    "bufio"
    "os"
    "strings"
    "fmt"
)

var inputReader *bufio.Reader

type Cfunc func() (bool, error)

func init() {
    inputReader = bufio.NewReader(os.Stdin)
}

func CRead() string {
    input,_ := inputReader.ReadString('\n') 
    return strings.TrimSpace(strings.TrimSuffix(input, "\n"))
}

// 格式化输出指令
func Coper(operate []string) {
    for k,v:=range operate {
        fmt.Printf("(%d) : %s\n", k,v)
    }
}

func CReturn(a Cfunc) bool {
    fmt.Println("============strart==========")
    b, err := a() 
    if err != nil {
        fmt.Println("错误信息：",err)
    }
    fmt.Println("===========end=============")

    return b

}


