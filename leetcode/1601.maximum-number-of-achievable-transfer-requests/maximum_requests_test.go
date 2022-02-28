package leetcode

import "testing"

func Test_maximumRequests(t *testing.T) {
	type args struct {
		n        int
		requests [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMaximumRequests01",
			args: args{n: 5, requests: [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}},
			want: 5,
		},
		{
			name: "testMaximumRequests02",
			args: args{n: 3, requests: [][]int{{0, 0}, {1, 2}, {2, 1}}},
			want: 3,
		},
		{
			name: "testMaximumRequests03",
			args: args{n: 4, requests: [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRequests(tt.args.n, tt.args.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
