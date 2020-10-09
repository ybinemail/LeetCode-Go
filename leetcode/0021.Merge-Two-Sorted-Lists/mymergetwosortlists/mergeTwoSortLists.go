package mymergetwosortlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 添加辅助头结点
	head := &ListNode{}
	head.Next = l1
	p := head
	q := l2
	for p.Next != nil {
		// 比较最小
		if q != nil && q.Val <= p.Next.Val {
			tmp := q
			q = q.Next
			tmp.Next = p.Next
			p.Next = tmp
		} else {
			p = p.Next
		}
	}
	// 仍然有遗留
	if q != nil {
		p.Next = q
	}

	return head.Next
}
