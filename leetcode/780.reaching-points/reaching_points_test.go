package leetcode

import "testing"

func Test_reachingPoints(t *testing.T) {
	type args struct {
		sx int
		sy int
		tx int
		ty int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testReachingPoints01",
			args: args{sx: 1, sy: 1, tx: 3, ty: 5},
			want: true,
		},
		{
			name: "testReachingPoints02",
			args: args{sx: 1, sy: 1, tx: 2, ty: 2},
			want: false,
		},
		{
			name: "testReachingPoints03",
			args: args{sx: 1, sy: 1, tx: 1, ty: 1},
			want: true,
		},
		{
			name: "testReachingPoints04",
			args: args{sx: 2, sy: 2, tx: 1, ty: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reachingPoints(tt.args.sx, tt.args.sy, tt.args.tx, tt.args.ty); got != tt.want {
				t.Errorf("reachingPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
