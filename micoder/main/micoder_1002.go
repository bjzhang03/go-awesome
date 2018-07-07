package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(line string) string {
	// 在此处理单行数据
	array:=strings.Split(line," ");

	result :=0;

	for _,v :=range array {
		tmp,_:=strconv.Atoi(v);
		result=result^tmp;
	}
	// 返回处理后的结果
	return strconv.Itoa(result);
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}