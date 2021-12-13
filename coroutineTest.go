package main

import (
	"net/http"
)

func testIo() {
	println("in testIo")

	go func() {
		println(" start baidu ")
		resp, err := http.Get("http://go.dev/")
		println("baidu status code = ",resp.StatusCode)
		if err != nil {
			println("error is ", err.Error())
		}
	}()

	go func() {
		println(" start google ")
		resp, err := http.Get("http://google.com")
		println("google status code = ",resp.StatusCode)
		//defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
		if err != nil {
			println("error is ", err.Error())
		}

	}()
}
