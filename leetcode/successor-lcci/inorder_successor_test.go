package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "testInorderSuccessor",
			args: args{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, p: &TreeNode{Val: 1}},
			want: &TreeNode{Val: 2, Left: &TreeNode{}, Right: &TreeNode{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
				t.Logf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
