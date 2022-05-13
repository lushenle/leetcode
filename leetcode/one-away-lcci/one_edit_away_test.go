package leetcode

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testOneEditAway01",
			args: args{first: "pale", second: "ple"},
			want: true,
		},
		{
			name: "testOneEditAway02",
			args: args{first: "pales", second: "pal"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
