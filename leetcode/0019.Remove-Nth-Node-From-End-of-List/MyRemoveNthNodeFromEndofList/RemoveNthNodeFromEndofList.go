package MyRemoveNthNodeFromEndofList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head
	p := tmp
	q := tmp

	i := 0
	for i <= n {
		q = q.Next
		i++
	}

	for q != nil {
		q = q.Next
		p = p.Next
	}
	p.Next = p.Next.Next
	return tmp.Next
}
