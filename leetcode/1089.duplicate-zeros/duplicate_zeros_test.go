package leetcode

import (
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test_duplicateZeros01",
			args: args{[]int{1, 2, 0, 3, 0, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
		})
	}
}
