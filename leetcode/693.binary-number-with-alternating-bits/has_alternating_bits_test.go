package leetcode

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testHasAlternatingBits01",
			args: args{n: 5},
			want: true,
		},
		{
			name: "testHasAlternatingBits02",
			args: args{n: 7},
			want: false,
		},
		{
			name: "testHasAlternatingBits03",
			args: args{n: 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
