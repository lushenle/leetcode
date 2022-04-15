package leetcode

import "testing"

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
			args: args{a: 10, b: 100},
			want: 100,
		},
		{
			name: "testMax02",
			args: args{a: -10, b: -100},
			want: -10,
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

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testRob01",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "testRob02",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
