package leetcode

import (
	"reflect"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test_minimumAbsDifference01",
			args: args{arr: []int{4, 2, 1, 3}},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			name: "test_minimumAbsDifference02",
			args: args{arr: []int{1, 3, 6, 10, 15}},
			want: [][]int{{1, 3}},
		},
		{
			name: "test_minimumAbsDifference03",
			args: args{arr: []int{3, 8, -10, 23, 19, -4, -14, 27}},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
