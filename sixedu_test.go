package sixedu

import (
    "testing"
    "fmt"
    "encoding/json"
)

type myFun func()

func TestFu(t *testing.T) {
    // fff := sum
    // fmt.Println("fff:",fff())
    newFuncoo := myFun(dispatch)
   // fmt.Println("newFuncoo:",newFuncoo.call())
    newFuncoo.call()
}

func (m myFun) call() {
   // fmt.Println("=====strt======")
   // dispatch()
   // fmt.Println("=====end======")
}

func sum() int {
    return 9
}

func dispatch() {
    //fmt.Println("disptch 111")
}

type Config struct {
    BasePath string `json:"base_path"`
    DataPath string `json:"data_path"`

}
func TestJson(t *testing.T) {
    c := Config{}
    cj1, _ := json.Marshal(c)
    fmt.Println(string(cj1))
    cj2 ,_ := json.MarshalIndent(c,"","  ")
    fmt.Println(string(cj2))

    js := `{
        "base_path":"这是路径",
        "data_path":"项目路径"
    }`
    cb := []byte(js)
    var c1 Config
    json.Unmarshal(cb, &c1)
    fmt.Println(c1)

}