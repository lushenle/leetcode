package leetcode

import (
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_wiggleSort01",
			args: args{nums: []int{1, 5, 1, 1, 6, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
		})
	}
}
