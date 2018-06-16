package main

import "fmt"

func longestPalindrome(s string) string {
	result := ""
	//把字符串变成长度为奇数的数据
	substr := make([]string, 0)
	for i := 0; i < len(s)-1; i++ {
		substr = append(substr, string(s[i]))
		substr = append(substr, "#")
	}
	substr = append(substr, string(s[len(s)-1]))
	fmt.Printf("substr = %s \n",substr)

	maxr := 0
	maxindex := 0
	for index, _ := range substr {
		r := 0
		left := index
		right := index
		for {
			if left < 0 || right >= len(substr) {
				break
			}
			if substr[left] == substr[right] {
				r++
				left--
				right++
			} else {
				break
			}
			if r > maxr {
				maxr = r
				maxindex = index
			}
		}
	}

	//fmt.Printf("maxr = %d, maxindex = %d\n",maxr,maxindex)
	for index, str := range substr {

		if index > (maxindex-maxr) && index < (maxindex+maxr) && str != "#" {
			result += str
		}
	}

	return result
}
func main() {
	fmt.Println(longestPalindrome("babad"))

}
