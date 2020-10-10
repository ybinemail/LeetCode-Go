package selectionSort

import (
	"testing"
)

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1",
			args{[]int{2, 6, 1, 8, 9, 5}},
			nil,
		},
		{"case2",
			args{[]int{9, 8, 7, 7, 1, 5}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectionSort(tt.args.arr)
			for _, v := range got {
				t.Logf("-> %d", v)
			}
			/*if got := bubbleSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
