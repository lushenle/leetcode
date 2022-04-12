package leetcode

import (
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	type args struct {
		widths []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testNumberOfLines01",
			args: args{widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, s: "abcdefghijklmnopqrstuvwxyz"},
			want: []int{3, 60},
		},
		{
			name: "testNumberOfLines02",
			args: args{widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, s: "bbbcccdddaaa"},
			want: []int{2, 4},
		},
		{
			name: "testNumberOfLines02",
			args: args{widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, s: ""},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLines(tt.args.widths, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
