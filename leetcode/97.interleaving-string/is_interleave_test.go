package leetcode

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_isInterleave01",
			args{s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac"},
			true,
		},
		{
			"test_isInterleave02",
			args{s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc"},
			false,
		},
		{
			"test_isInterleave03",
			args{s1: "", s2: "", s3: ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
