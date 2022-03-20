package leetcode

import "testing"

func Test_networkBecomesIdle(t *testing.T) {
	type args struct {
		edges    [][]int
		patience []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testNetworkBecomesIdle01",
			args: args{edges: [][]int{{0, 1}, {1, 2}}, patience: []int{0, 2, 1}},
			want: 8,
		},
		{
			name: "testNetworkBecomesIdle02",
			args: args{edges: [][]int{{0, 1}, {0, 2}, {1, 2}}, patience: []int{0, 10, 10}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkBecomesIdle(tt.args.edges, tt.args.patience); got != tt.want {
				t.Errorf("networkBecomesIdle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMax01",
			args: args{a: 10, b: 9},
			want: 10,
		},
		{
			name: "testMax02",
			args: args{a: 8, b: 9},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
