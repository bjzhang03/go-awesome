package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	result := 0
	current := make(map[uint8]int)
	//使用index来标记开始位置的数据
	index := -1
	j := 0
	for ; j < len(s); j++ {
		//判断数据是否出现过了
		if _, ok := current[s[j]]; ok {
			//因为重复出现了，需要判断一下，重复出现的之间的字符串的长度
			if result < j-index-1 {
				result = j - index - 1
			}
			if index < current[s[j]] {
				//如果当前的结点已经不是第一个的开始节点的时候，更新开始节点的位置
				index = current[s[j]]
			}
			//如果当前的位置已经在前面出现了，则更新数据
			current[s[j]] = j
		} else {
			//数据以前没有出现过
			current[s[j]] = j
		}
	}
	//处理最后一个可能出现的最长字符串的数据
	if result < j-index-1 {
		result = j - index - 1
	}
	return result
}

//success
func main() {
	s := "abcabcbb"
	//s := "dda"
	//s := "abba"
	//s := "abcb"
	fmt.Println(lengthOfLongestSubstring(s))

}
