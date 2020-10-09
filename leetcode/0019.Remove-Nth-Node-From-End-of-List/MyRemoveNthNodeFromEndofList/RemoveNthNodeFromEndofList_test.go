package MyRemoveNthNodeFromEndofList

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {

	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case_1", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 1}, nil},
		{"case_2", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 2}, nil},
		{"case_3", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 4}, nil},
		{"case_4", args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 3}, nil},
		{"case_5", args{&ListNode{1, &ListNode{2, nil}}, 1}, nil},
		{"case_6", args{&ListNode{1, &ListNode{2, nil}}, 2}, nil},
		{"case_7", args{&ListNode{1, nil}, 1}, nil},
		{"case_8", args{&ListNode{}, 1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			for {
				if got != nil {
					t.Logf("-> %d", got.Val)
					got = got.Next
				} else {
					break
				}
			}
			//if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			//}
		})
	}
}
