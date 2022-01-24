package sort

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
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
				arr: []int{3, 7, 1, 2, 8, 4},
			},
			want: []int{1, 2, 3, 4, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}
