package leetcode

import "testing"

func Test_countValidWords(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{sentence: "cat and  dog"},
			want: 3,
		},
		{
			name: "test02",
			args: args{sentence: "!obama 1hillary "},
			want: 0,
		},
		{
			name: "test03",
			args: args{sentence: "he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."},
			want: 6,
		},
		{
			name: "test04",
			args: args{sentence: "b-a-c f-d"},
			want: 1,
		},
		{
			name: "test04",
			args: args{sentence: "f-d."},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countValidWords(tt.args.sentence); got != tt.want {
				t.Errorf("countValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
