package leetcode

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{haystack: "adele", needle: "le"},
			want: 3,
		},
		{
			name: "test02",
			args: args{haystack: "", needle: "e"},
			want: -1,
		},
		{
			name: "test03",
			args: args{haystack: "adele", needle: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
