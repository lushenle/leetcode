package leetcode

import "testing"

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_minOperations01",
			args: args{nums: []int{1, 1, 1}},
			want: 3,
		},
		{
			name: "test_minOperations02",
			args: args{nums: []int{1, 5, 2, 4, 1}},
			want: 14,
		},
		{
			name: "test_minOperations03",
			args: args{nums: []int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
