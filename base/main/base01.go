package main

import "fmt"

func main()  {
	first := "fisrt"
	fmt.Println([]rune(first))
	fmt.Println([]byte(first))

	second := "社区"
	fmt.Println([]rune(second))
	fmt.Println([]byte(second))

	s1 := "golangcaff"
	fmt.Println(s1[:3])
	s2 := "截取中文"
	//试试这样能不能截取?
	fmt.Println(s2[:2])
	s3 := "截取中文"
	//试试这样能不能截取?
	res := []rune(s3)
	fmt.Println(string(res[:2]))

	s4 := "截取中文"
	//试试这样能不能截取?
	fmt.Println(s4[:3])

}
