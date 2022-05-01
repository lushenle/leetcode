package leetcode

import (
	"reflect"
	"testing"
)

func Test_getAllElements(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testGetAllElements01",
			args: args{
				root1: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}},
				root2: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 3}}},
			want: []int{0, 1, 1, 2, 3, 4},
		},
		{
			name: "testGetAllElements02",
			args: args{
				root1: &TreeNode{Val: 1, Right: &TreeNode{Val: 8}},
				root2: &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}},
			want: []int{1, 1, 8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
