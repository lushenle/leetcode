package leetcode

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testDiStringMatch01",
			args: args{s: "IDID"},
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "testDiStringMatch02",
			args: args{s: "III"},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "testDiStringMatch02",
			args: args{s: "DDI"},
			want: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
