package leetcode

import "testing"

func Test_cutOffTree(t *testing.T) {
	type args struct {
		forest [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCutOffTree01",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}}},
			want: 6,
		},
		{
			name: "testCutOffTree02",
			args: args{forest: [][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}}},
			want: -1,
		},
		{
			name: "testCutOffTree03",
			args: args{forest: [][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutOffTree(tt.args.forest); got != tt.want {
				t.Errorf("cutOffTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
