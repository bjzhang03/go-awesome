package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

func main() {
	body := "{\"action\":20}"
	res, err := http.Post("http://www.baidu.com", "application/json;charset=utf-8", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}