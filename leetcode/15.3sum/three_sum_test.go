package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testThreeSum01",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "testThreeSum02",
			args: args{nums: []int{}},
			want: [][]int{},
		},
		{
			name: "testThreeSum02",
			args: args{nums: []int{0}},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_next(t *testing.T) {
	type args struct {
		nums []int
		l    int
		r    int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "testNext01",
			args:  args{nums: []int{1, 2, 3, 4}, l: 1, r: 2},
			want:  2,
			want1: 1,
		},
		{
			name:  "testNext02",
			args:  args{nums: []int{1, 2, 3, 4}, l: 3, r: 2},
			want:  3,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := next(tt.args.nums, tt.args.l, tt.args.r)
			if got != tt.want {
				t.Errorf("next() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("next() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
