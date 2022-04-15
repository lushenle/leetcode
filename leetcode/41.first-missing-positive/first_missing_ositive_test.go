package leetcode

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFirstMissingPositive01",
			args: args{nums: []int{1, 2, 0}},
			want: 3,
		},
		{
			name: "testFirstMissingPositive02",
			args: args{nums: []int{3, 4, -1, 1}},
			want: 2,
		},
		{
			name: "testFirstMissingPositive02",
			args: args{nums: []int{7, 8, 9, 11}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive1(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
