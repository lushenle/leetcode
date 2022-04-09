package leetcode

import "testing"

func Test_uniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testUniqueMorseRepresentations01",
			args: args{words: []string{"gin", "zen", "gig", "msg"}},
			want: 2,
		},
		{
			name: "testUniqueMorseRepresentations02",
			args: args{words: []string{"a"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.args.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
