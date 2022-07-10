package leetcode

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_minCostClimbingStairs01",
			args{cost: []int{10, 15, 20}},
			15,
		},
		{
			"test_minCostClimbingStairs02",
			args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
