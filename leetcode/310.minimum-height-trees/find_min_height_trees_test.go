package leetcode

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testFindMinHeightTrees01",
			args: args{n: 4, edges: [][]int{{1, 0}, {1, 2}, {1, 3}}},
			want: []int{1},
		},
		{
			name: "testFindMinHeightTrees02",
			args: args{n: 6, edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
