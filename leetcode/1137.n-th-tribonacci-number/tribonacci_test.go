package leetcode

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{n: 4},
			want: 4,
		},
		{
			name: "test02",
			args: args{n: 6},
			want: 13,
		},
		{
			name: "test03",
			args: args{n: 25},
			want: 1389537,
		},
		{
			name: "test04",
			args: args{n: 37},
			want: 2082876103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci1(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
