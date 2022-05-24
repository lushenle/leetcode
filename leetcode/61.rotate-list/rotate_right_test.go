package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testRotateRight01",
			args: args{head: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5}}}}}, k: 2},
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5,
				Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			name: "testRotateRight02",
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, k: 4},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			name: "testRotateRight03",
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, k: 3},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		},
		{
			name: "testRotateRight04",
			args: args{head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, k: 5},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
