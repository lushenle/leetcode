package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMyAtoi01",
			args: args{s: "42"},
			want: 42,
		},
		{
			name: "testMyAtoi02",
			args: args{s: "    -42"},
			want: -42,
		},
		{
			name: "testMyAtoi03",
			args: args{s: "4193 with words"},
			want: 4193,
		},
		{
			name: "testMyAtoi04",
			args: args{s: "-91283472332"},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
