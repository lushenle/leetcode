package leetcode

import "testing"

func Test_findClosest(t *testing.T) {
	type args struct {
		words []string
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFindClosest",
			args: args{words: []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}, word1: "a", word2: "student"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosest(tt.args.words, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
