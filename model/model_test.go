package model

import (
	"testing"
    "fmt"
)

func TestRfdata(t *testing.T) {
	userDatas := make(map[string]Model, 0)
	rfdata("users", "username", userDatas)

    for _,v := range userDatas {
        fmt.Println(v.ToString())
    }

	t.Log(userDatas)


}
