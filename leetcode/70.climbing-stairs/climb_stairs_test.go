package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	res := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597}
	for k, v := range res {
		tests := []struct {
			name string
			args args
			want int
		}{
			{
				name: "testN",
				args: args{n: k + 1},
				want: v,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := climbStairs(tt.args.n); got != tt.want {
					t.Errorf("climbStairs() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
