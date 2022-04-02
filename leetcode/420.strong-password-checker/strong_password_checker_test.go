package leetcode

import "testing"

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testStrongPasswordChecker01",
			args: args{password: "a"},
			want: 5,
		},
		{
			name: "testStrongPasswordChecker02",
			args: args{password: "aA1"},
			want: 3,
		},
		{
			name: "testStrongPasswordChecker03",
			args: args{password: "1337C0d3"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPasswordChecker(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
