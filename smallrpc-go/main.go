package main

import (
	"context"
	"encoding/json"
	"time"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsontest() {
	stu := Student{
		Name: "xsax",
		Age:  15,
	}
	jsonbyte, _ := json.Marshal(stu) // 序列化为json字符串
	println(string(jsonbyte))
	jsonstr := "{\"name\":\"xsax\",\"age\":15}"

	stu2 := Student{}
	_ = json.Unmarshal([]byte(jsonstr), &stu2) // 反序列化
	println(stu2.Name)
	println(stu2.Age)
}

func main() {
	ctx, done := context.WithTimeout(context.Background(), time.Second*2)

	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				println("----------")
				return
			default:
				println("111111111111111111")
			}
		}
	}(ctx)

	time.Sleep(time.Second * 5)
	done()
}
