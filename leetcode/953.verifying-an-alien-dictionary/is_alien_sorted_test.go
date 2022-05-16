package leetcode

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testIsAlienSorted01",
			args: args{words: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz"},
			want: true,
		},
		{
			name: "testIsAlienSorted02",
			args: args{words: []string{"word", "world", "row"}, order: "worldabcefghijkmnpqstuvxyz"},
			want: false,
		},
		{
			name: "testIsAlienSorted03",
			args: args{words: []string{"apple", "app"}, order: "abcdefghijklmnopqrstuvwxyz"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
