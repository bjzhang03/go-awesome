package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(line string) string {
	// 在此处理单行数据
	count,_:=strconv.Atoi(line)
	var result []int =make ([]int,count+1)
	result[0]=1;
	result[1]=1;
	for i:=2;i<=count;i++{
		result[i]=result[i-1]+result[i-2]
	}

	// 返回处理后的结果
	return strconv.Itoa(result[count])
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}