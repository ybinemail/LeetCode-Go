package mymergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	} else if length < 2 {
		return lists[0]
	}
	middle := length / 2
	left := lists[0:middle]
	right := lists[middle:]
	return mergeTwoLists(mergeKLists(left), mergeKLists(right))
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
