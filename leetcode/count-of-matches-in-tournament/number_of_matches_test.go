package leetcode

import (
	"math/rand"
	"testing"
	"time"
)

func Test_numberOfMatches(t *testing.T) {
	type args struct {
		n int
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(1_000-1) + 1
		tests := []struct {
			name string
			args args
			want int
		}{
			{
				name: "testN",
				args: args{n: n},
				want: n - 1,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := numberOfMatches(tt.args.n); got != tt.want {
					t.Errorf("numberOfMatches() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
