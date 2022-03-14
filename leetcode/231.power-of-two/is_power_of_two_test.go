package leetcode

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testIsPowerOfTwo01",
			args: args{n: 8},
			want: true,
		},
		{
			name: "testIsPowerOfTwo02",
			args: args{n: 7},
			want: false,
		},
		{
			name: "testIsPowerOfTwo02",
			args: args{n: -167},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
