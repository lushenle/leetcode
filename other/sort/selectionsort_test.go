package sort

import (
	"reflect"
	"testing"
)

func Test_findSmallest(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{
				arr: []int{3, 6, 1, 9, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallest(tt.args.arr); got != tt.want {
				t.Errorf("findSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{
				arr: []int{5, 2, 3, 7, 9, 1},
			},
			want: []int{1, 2, 3, 5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
