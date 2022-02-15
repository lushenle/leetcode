package leetcode

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{digits: []int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
		{
			name: "test02",
			args: args{digits: []int{0}},
			want: []int{1},
		},
		{
			name: "test03",
			args: args{digits: []int{1, 2, 9, 2}},
			want: []int{1, 2, 9, 3},
		},
		{
			name: "test04",
			args: args{digits: []int{129}},
			want: []int{130},
		},
		{
			name: "test05",
			args: args{digits: []int{99}},
			want: []int{100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
