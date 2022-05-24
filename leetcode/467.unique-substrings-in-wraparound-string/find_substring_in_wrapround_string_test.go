package leetcode

import "testing"

func Test_findSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFindSubstringInWraproundString01",
			args: args{p: "a"},
			want: 1,
		},
		{
			name: "testFindSubstringInWraproundString02",
			args: args{p: "cac"},
			want: 2,
		},
		{
			name: "testFindSubstringInWraproundString03",
			args: args{p: "zab"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstringInWraproundString(tt.args.p); got != tt.want {
				t.Errorf("findSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}
