package main

import (
	"fmt"
	"strconv"
)

//新建一个listnode的类型数据
type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	result := ""
	for it := p; it != nil; it = it.Next {
		result = result + strconv.Itoa(it.Val) + " "

	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//新建结果数据
	result := &ListNode{0, nil}
	it := result
	carry := 0

	for ; l1 != nil && l2 != nil; {
		//计算当前的值
		temp := &ListNode{(l1.Val + l2.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l1.Val + l2.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	//l1没有到nil
	for ; l1 != nil; {
		//计算当前的值
		temp := &ListNode{(l1.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l1.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l1 = l1.Next
	}
	//l2没有到nil
	for ; l2 != nil; {
		//计算当前的值
		temp := &ListNode{(l2.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l2.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l2 = l2.Next
	}
	//carry还有数据
	if carry > 0 {
		temp := &ListNode{carry, nil}
		it.Next = temp
		//计算进位数据
	}

	return result.Next
}
//success
func main() {
	var l1 *ListNode
	var l2 *ListNode

	l1 = &ListNode{9, nil}
	header1 := l1
	//l1.Next = &ListNode{4, nil}
	//l1 = l1.Next
	//l1.Next = &ListNode{3, nil}

	l2 = &ListNode{9, nil}
	header2 := l2
	//l2.Next = &ListNode{6, nil}
	//l2 = l2.Next
	//l2.Next = &ListNode{4, nil}

	for it := header1; it != nil; it = it.Next {
		fmt.Printf("%d -> ", it.Val)
	}

	fmt.Println(addTwoNumbers(header1, header2))
}
