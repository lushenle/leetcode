package leetcode

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		mis []int
		njs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testCombine01",
			args: args{mis: []int{1, 3}, njs: []int{2}},
			want: []int{1, 2, 3},
		},
		{
			name: "testCombine02",
			args: args{mis: []int{1, 2}, njs: []int{3, 4}},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.mis, tt.args.njs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testFindMedianSortedArrays01",
			args: args{nums1: []int{1, 3}, nums2: []int{2}},
			want: 2.0,
		},
		{
			name: "testFindMedianSortedArrays02",
			args: args{nums1: []int{1, 2}, nums2: []int{3, 4}},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_medianOf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testMedianOf01",
			args: args{nums: []int{1, 2, 3}},
			want: 2.0,
		},
		{
			name: "testMedianOf02",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianOf(tt.args.nums); got != tt.want {
				t.Errorf("medianOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
