package leetcode

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testLongestCommonSubsequence01",
			args: args{text1: "abcde", text2: "ace"},
			want: 3,
		},
		{
			name: "testLongestCommonSubsequence02",
			args: args{text1: "abc", text2: "abc"},
			want: 3,
		},
		{
			name: "testLongestCommonSubsequence03",
			args: args{text1: "abc", text2: "def"},
			want: 0,
		},
		{
			name: "testLongestCommonSubsequence04",
			args: args{text1: "abcdef", text2: "ace"},
			want: 3,
		},
		{
			name: "testLongestCommonSubsequence05",
			args: args{text1: "ezupkr", text2: "ubmrapg"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
