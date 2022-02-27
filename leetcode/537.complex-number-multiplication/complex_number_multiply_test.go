package leetcode

import "testing"

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testComplexNumberMultiply01",
			args: args{num1: "1+1i", num2: "1+1i"},
			want: "0+2i",
		},
		{
			name: "testComplexNumberMultiply02",
			args: args{num1: "1+-1i", num2: "1+-1i"},
			want: "0+-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply1(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "testParse01",
			args:  args{s: "1+-1i"},
			want:  1,
			want1: -1,
		},
		{
			name:  "testParse01",
			args:  args{s: "1+1i"},
			want:  1,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.args.s)
			if got != tt.want {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
