package leetcode

import (
	"reflect"
	"testing"
)

func Test_kthPalindrome(t *testing.T) {
	type args struct {
		queries   []int
		intLength int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "test_kthPalindrome01",
			args: args{queries: []int{1, 2, 3, 4, 5}, intLength: 1},
			want: []int64{1, 2, 3, 4, 5},
		},
		{
			name: "test_kthPalindrome02",
			args: args{queries: []int{1, 2, 3, 4, 5, 90}, intLength: 3},
			want: []int64{101, 111, 121, 131, 141, 999},
		},
		{
			name: "test_kthPalindrome03",
			args: args{queries: []int{2, 201429812, 8, 520498110, 492711727, 339882032, 462074369, 9, 7, 6}, intLength: 1},
			want: []int64{2, -1, 8, -1, -1, -1, -1, 9, 7, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthPalindrome(tt.args.queries, tt.args.intLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
