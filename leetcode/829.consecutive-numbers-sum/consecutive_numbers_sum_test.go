package leetcode

import "testing"

func Test_consecutiveNumbersSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_consecutiveNumbersSum01",
			args: args{n: 11},
			want: 2,
		},
		{
			name: "test_consecutiveNumbersSum02",
			args: args{n: 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.n); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
