package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "test02",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "test03",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: "test04",
			args: args{s: "Today is a nice day"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
