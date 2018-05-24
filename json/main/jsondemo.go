package main

import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

func main() {
	//构建一个对象
	st := &Student{
		"TestData",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	fmt.Println("before JSON encoding :")
	//对对象进行显示
	st.ShowStu()

	//对对象进行json转化
	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		//以字符串的方式输出转化之后的json格式的数据
		fmt.Println(string(b))
	}

	//使用一个channel进行传输
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch

	//传输结束，接受到数据进行进一步的处理操作
	fmt.Println("--------------------------------")
	stb := &Student{}
	stb.ShowStu()
	//将结构体从json中还原出来
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
}
