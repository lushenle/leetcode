package search

import (
	"math/rand"
	"testing"
	"time"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}

	//arr := make([]int, 100)
	var arr []int
	for i := 1; i < 1_000; i++ {
		arr = append(arr, i)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		v := rand.Intn(1_000-1) + 1
		tests := []struct {
			name string
			args args
			want int
		}{
			{
				name: "testN",
				args: args{arr: arr, target: v},
				want: v - 1,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
					t.Errorf("binarySearch() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
