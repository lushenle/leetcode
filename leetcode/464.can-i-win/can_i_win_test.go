package leetcode

import "testing"

func Test_canIWin(t *testing.T) {
	type args struct {
		maxChoosableInteger int
		desiredTotal        int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testCanIWin01",
			args: args{maxChoosableInteger: 10, desiredTotal: 11},
			want: false,
		},
		{
			name: "testCanIWin02",
			args: args{maxChoosableInteger: 10, desiredTotal: 1},
			want: true,
		},
		{
			name: "testCanIWin03",
			args: args{maxChoosableInteger: 10, desiredTotal: 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canIWin(tt.args.maxChoosableInteger, tt.args.desiredTotal); got != tt.want {
				t.Errorf("canIWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
