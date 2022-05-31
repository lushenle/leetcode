package leetcode

import "testing"

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_alienOrder01",
			args: args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want: "wertf",
		},
		{
			name: "test_alienOrder02",
			args: args{words: []string{"z", "x"}},
			want: "zx",
		},
		{
			name: "test_alienOrder03",
			args: args{words: []string{"z", "x", "z"}},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
