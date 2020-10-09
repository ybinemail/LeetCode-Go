package mymergetwosortlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	head.Next = l1
	p := head
	q := l2
	for p.Next != nil {
		if q != nil && q.Val <= p.Next.Val {
			tmp := q
			q = q.Next
			tmp.Next = p.Next
			p.Next = tmp
		} else {
			p = p.Next
		}
	}
	if q != nil {
		p.Next = q
	}

	return head.Next
}
