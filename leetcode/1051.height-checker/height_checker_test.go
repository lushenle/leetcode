package leetcode

import "testing"

func Test_heightChecker(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_heightChecker01",
			args: args{heights: []int{1, 1, 4, 2, 1, 3}},
			want: 3,
		},
		{
			name: "test_heightChecker02",
			args: args{heights: []int{5, 1, 2, 3, 4}},
			want: 5,
		},
		{
			name: "test_heightChecker03",
			args: args{heights: []int{1, 2, 3, 4, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heightChecker(tt.args.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
