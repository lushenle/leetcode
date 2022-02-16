package leetcode

import "testing"

func Test_checkWays(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{[][]int{{1, 2}, {2, 3}}},
			want: 1,
		},
		{
			name: "test02",
			args: args{[][]int{{1, 2}, {2, 3}, {1, 3}}},
			want: 2,
		},
		{
			name: "test03",
			args: args{[][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkWays(tt.args.pairs); got != tt.want {
				t.Errorf("checkWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
