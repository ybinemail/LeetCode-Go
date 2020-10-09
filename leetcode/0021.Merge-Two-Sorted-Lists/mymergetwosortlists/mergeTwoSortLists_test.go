package mymergetwosortlists

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{&ListNode{1, &ListNode{3, &ListNode{5, nil}}}, &ListNode{2, &ListNode{6, nil}}}, nil},
		{"case2", args{&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4,nil}}}}, nil},
		{"case3", args{&ListNode{1, nil}, &ListNode{1, &ListNode{3, &ListNode{4,nil}}}}, nil},
		{"case3", args{&ListNode{}, &ListNode{1, &ListNode{3, &ListNode{4,nil}}}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.l1, tt.args.l2)
			for got != nil {
				t.Logf("-> %d", got.Val)
				got = got.Next

			}
			/*if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
