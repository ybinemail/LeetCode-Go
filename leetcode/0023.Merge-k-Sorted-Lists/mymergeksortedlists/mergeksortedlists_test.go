package mymergeksortedlists

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"case1", args{
			[]*ListNode{
				&ListNode{1, &ListNode{4, &ListNode{5, nil}}},
				&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
				&ListNode{2, &ListNode{6, nil}},
			},
		}, nil},
		{"case2", args{
			nil,
		}, nil},
		{"case3", args{
			[]*ListNode{
				nil,
			},
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			for got != nil {
				t.Logf("-> %d", got.Val)
				got = got.Next

			}
			/*if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
