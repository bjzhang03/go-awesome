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
	lineArr := strings.Split(line, " ")
	a, _ := strconv.Atoi(lineArr[0])
	b, _ := strconv.Atoi(lineArr[1])
	// 返回处理后的结果
	return strconv.Itoa(a + b)
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
