package leetcode

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFindMin01",
			args: args{nums: []int{4, 5, 6, 1, 2, 3}},
			want: 1,
		},
		{
			name: "testFindMin02",
			args: args{nums: []int{-4, -5, -6, -1, -2, -3}},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
