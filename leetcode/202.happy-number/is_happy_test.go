package leetcode

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testIsHappy01",
			args: args{n: 2},
			want: false,
		},
		{
			name: "testIsHappy02",
			args: args{n: 19},
			want: true,
		},
		{
			name: "testIsHappy03",
			args: args{n: 347},
			want: false,
		},
		{
			name: "testIsHappy04",
			args: args{n: 13},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
