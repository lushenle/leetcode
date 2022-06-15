package leetcode

import "testing"

func Test_smallestDistancePair(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_smallestDistancePair01",
			args: args{nums: []int{1, 3, 1}, k: 1},
			want: 0,
		},
		{
			name: "test_smallestDistancePair02",
			args: args{nums: []int{1, 1, 1}, k: 2},
			want: 0,
		},
		{
			name: "test_smallestDistancePair02",
			args: args{nums: []int{1, 6, 1}, k: 3},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDistancePair(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
