package MyRemoveNthNodeFromEndofList

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	/*p6 := ListNode{6, nil}
	p5 := ListNode{5, &p6}
	p4 := ListNode{4, &p5}
	p3 := ListNode{3, &p4}*/
	p2 := ListNode{2, nil}
	p1 := ListNode{1, &p2}
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case_3", args{&p1, 1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
