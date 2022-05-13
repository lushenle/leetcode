package leetcode

import (
	"testing"
)

func Test_minStickers(t *testing.T) {
	type args struct {
		stickers []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMinStickers01",
			args: args{stickers: []string{"with", "example", "science"}, target: "thehat"},
			want: 3,
		},
		{
			name: "testMinStickers02",
			args: args{stickers: []string{"notice", "possible"}, target: "basicbasic"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStickers(tt.args.stickers, tt.args.target); got != tt.want {
				t.Errorf("minStickers() = %v, want %v", got, tt.want)
			}
		})
	}
}
