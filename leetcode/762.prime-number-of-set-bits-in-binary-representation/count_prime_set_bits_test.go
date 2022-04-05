package leetcode

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testCountPrimeSetBits01",
			args: args{left: 6, right: 10},
			want: 4,
		},
		{
			name: "testCountPrimeSetBits02",
			args: args{left: 10, right: 15},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
