package main

import (
	"example.com/greetings"
	"fmt"
	"helloxyy.com/mylove-go/model"
	"helloxyy.com/utilsmod"
	"time"
)

func main() {
	//resp, err := http.Get("http://google.com")
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//if err != nil {
	//	println("error is ", err.Error())
	//}
	//testHello()
	testIo()

	time.Sleep(100 * time.Second)
	fmt.Println("done")
}

func testHello() {
	var name = "lihao"
	fmt.Printf("Hello, World! %s \n", name)
	var ret = greetings.Hello("lihao baobao")
	fmt.Printf("result is %s\n", ret)

	fmt.Println(model.Eat("food"))

	fmt.Println(utils.LowerStr("xiaoxiao"))
}
