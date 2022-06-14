package leetcode

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_findDiagonalOrder01",
			args: args{mat: [][]int{{1, 2}, {3, 4}}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "test_findDiagonalOrder02",
			args: args{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "test_findDiagonalOrder02",
			args: args{mat: [][]int{{}}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
