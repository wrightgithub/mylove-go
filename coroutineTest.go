package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func testIo() {
	println("in testIo")

	go func() {
		println(" start baidu ")
		resp, err := http.Get("http://baidu.com/")
		if err != nil {
			println("error is ", err.Error())
			return
		}
		println("baidu status code = ",resp.StatusCode)
	}()

	//go func() {
	//	println(" start google ")
	//	resp, err := http.Get("http://google.com")
	//	println("google status code = ",resp.StatusCode)
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(body))
	//	if err != nil {
	//		println("error is ", err.Error())
	//	}
	//
	//}()
	//
	go func() {
		println(" start google timeout test")
		client := http.Client{Timeout: 10 * time.Second}
		resp, err :=client.Get("http://google.com");
		if err != nil {
			println("error is ", err.Error())
			return
		}
		println("google status code = ",resp.StatusCode)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

	}()
}
