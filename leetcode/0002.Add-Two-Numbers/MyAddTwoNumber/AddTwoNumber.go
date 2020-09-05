package MyAddTwoNumber

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 已p1为主,按照加法规则，进行相加处理
	p1 := l1
	p2 := l2
	for {
		p1.Val = p1.Val + p2.Val
		if p1.Next == nil || p2.Next == nil {
			break
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}

	if p2.Next != nil {
		p1.Next = p2.Next
	}
	addSum(l1, 0)
	return l1
}

func addSum(link *ListNode, add int) {
	p := link
	if p == nil {
		return
	}
	if (p.Val + add) >= 10 {
		p.Val = (p.Val + add) % 10
		add = 1
	} else {
		p.Val = p.Val + add
		add = 0
	}
	if p.Next == nil && add == 1 {
		p.Next = &ListNode{1, nil}
		add = 0
	}
	addSum(p.Next, add)
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 偷懒转换为int 无法针对超大数据，进行处理
	p1 := l1
	p2 := l2
	num1 := 0
	num2 := 0
	for i := 1; p1 != nil || p2 != nil; i = i * 10 {
		if p1 != nil {
			num1 = num1 + p1.Val*i
			p1 = p1.Next
		}
		if p2 != nil {
			num2 = num2 + p2.Val*i
			p2 = p2.Next
		}
	}
	sum := num1 + num2
	l := getLinkNode(strconv.Itoa(sum))
	return l
}

func getLinkNode(numstr string) *ListNode {
	link := &ListNode{}
	p := link
	for i := len(numstr) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(numstr[i : i+1])
		if err != nil {
			fmt.Printf("strconv error : %s", err)
		}
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	return link.Next
}
