package leetcode

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMaximumDifference01",
			args: args{nums: []int{7, 1, 5, 4}},
			want: 4,
		},
		{
			name: "testMaximumDifference02",
			args: args{nums: []int{9, 4, 3, 2}},
			want: -1,
		},
		{
			name: "testMaximumDifference03",
			args: args{nums: []int{1, 5, 2, 10}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
