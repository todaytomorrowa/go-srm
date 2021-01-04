package model

import (
    "os"
    "errors"
    "bufio"
    "io"
    "strings"
    "reflect"
    "strconv"
    "fmt"
   
)

type Model interface {
    ToString() string
}

var (
    path string = "/Users/derek/go/src/sixedu/data/" // 数据路径
    suffix string = ".sql" //文件后缀
    models map[string]interface{} //记录标示  例如：users，admin
)

func init() {
    models = make(map[string]interface{})
    models["users"] = NewUser 

    userDatas = make(map[string]Model, 0)
	rfdata("users", "username", userDatas)
}

func rfdata(name, primay string, datas map[string]Model) error{

    //fmt.Println(44)
    //1.文件中读取数据
    //2.遍历数据（字段，数据） 通过"\n"分割
        //2.1判断是否为字段
        //2.2存储数据到models
            //2.2.1 根据name 获取到模型
            //2.2.2 利用反射-》对模型增删改查操作
            //2.2.3 再把模型数据存入datas
    fileHundle, err := os.Open(path + name + suffix) //根据文件路径读取数据
    if err != nil {
        //fmt.Println("文件查询不到,",err)
        return errors.New("文件查询不到")
    }
    defer fileHundle.Close()

    buf := bufio.NewReader(fileHundle)

    field := make([]string,0)

    for {
        
        row , rerr := buf.ReadBytes('\n') //根据换行读取文件信息
        if rerr != nil {
            if rerr == io.EOF {
                break
            }
            fmt.Println("读取文件异常，",rerr)
        }
        

        data := strings.Split(string(row),",")

       // fmt.Println("获取数据：",data)


        if len(field) == 0 {
            field = data

            for k,v := range data {
                field[k] = strings.TrimSpace(strings.TrimSuffix(v , "\n")) 
            }

        } else {
            //2.2存储数据到models
                //2.2.1 根据name 获取到模型
                //2.2.2 利用反射-》对模型增删改查操作
                //2.2.3 再把模型数据存入datas
                toModel(name, primay, datas, data, field)
        }
        
    }
    //fmt.Println("获取字段值：",field)
    return nil
}

func toModel(name, primay string, datas map[string]Model,data,field []string) error {
    // 先判断模型是否存在
    if models[name] == nil {
        return errors.New("不存在模型："+name)
    }
    modelV := reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]
   // fmt.Printf("modelV type %T\n",modelV)

   // fmt.Printf("field : %v\n", field)

    var primayValue string
    for k, v := range data {
        if primay == field[k] {
            primayValue = v
           // fmt.Printf("获取到的主键值：%v\n", primayValue)
        }
       // fmt.Printf("k: %v, v:%v \n",k,v)
        fset := modelV.MethodByName("Set" + strings.Title(field[k]))
        fset.Call([]reflect.Value{
            reflect.ValueOf(toTypeValue(modelV, field[k], v)),
        })

        modelV.Elem().FieldByName(field[k]).Type().Name()

        //fmt.Println("mytype:",mytype)
    }

   // fmt.Println(datas)

        
    datas[primayValue] = modelV.Interface().(Model)

    return nil
}

func toTypeValue(modelV reflect.Value, field, value string) interface{} {
    mtype := modelV.Elem().FieldByName(field).Type().Name()
    switch mtype {
        case "int":
            b, _ := strconv.Atoi(value)
            return b
    }
    return string(value)

}

func fwrite(name string, models map[string]Model) bool {
    countent := getModelTostring(models)
    outfile, outerr := os.OpenFile(path+name+suffix,os.O_WRONLY|os.O_CREATE,0666)
    if outerr != nil {
        fmt.Println("文件打开失败")
        return false
    }
    defer outfile.Close()

    outbufwri := bufio.NewWriter(outfile)
    outbufwri.WriteString(countent + "\n")
    outbufwri.Flush()
    return true
}

func rwdata() {

}

func getModelTostring(models map[string]Model) string {
    var fields string 
    var countents string
    for _, model := range models {
        if fields == "" {
            remodel := reflect.ValueOf(model)
            modelZ := reflect.Zero(remodel.Type().Elem())
            for i:=0;i<modelZ.NumField();i++{
                fields = fields + modelZ.Type().Field(i).Name + ","
            }
            fields = strings.TrimSuffix(fields,",")
        }
        countents = countents + model.ToString() + "\n"
    }
    return fields + "\n" + strings.TrimSuffix(countents, "\n")
    
}