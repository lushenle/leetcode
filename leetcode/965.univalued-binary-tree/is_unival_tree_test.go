package leetcode

import "testing"

func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testIsUnivalTree01",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}}},
			want: true,
		},
		{
			name: "testIsUnivalTree02",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}},
			want: false,
		},
		{
			name: "testIsUnivalTree03",
			args: args{root: nil},
			want: true,
		},
		{
			name: "testIsUnivalTree04",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
