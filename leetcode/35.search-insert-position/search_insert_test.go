package leetcode

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{1, 3, 5, 6}, target: 5},
			want: 2,
		},
		{
			name: "test02",
			args: args{nums: []int{1, 3, 5, 6}, target: 4},
			want: 2,
		},
		{
			name: "test03",
			args: args{nums: []int{1, 3, 5, 6}, target: 0},
			want: 0,
		},
		{
			name: "test04",
			args: args{nums: []int{1, 3, 5, 6}, target: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
