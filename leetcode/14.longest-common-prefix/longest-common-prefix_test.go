package leetcode

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{
				strs: []string{"ab", "a"},
			},
			want: "a",
		},
		{
			name: "test02",
			args: args{
				strs: []string{"abb", "abc"},
			},
			want: "ab",
		},
		{
			name: "test03",
			args: args{
				strs: []string{"trumpa", "trumpb", "trump"},
			},
			want: "trump",
		},
		{
			name: "test04",
			args: args{
				strs: []string{"", "biden", ""},
			},
			want: "",
		},
		{
			name: "test06",
			args: args{
				strs: []string{"", "b"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
