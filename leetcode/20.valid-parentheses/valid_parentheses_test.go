package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				s: "(",
			},
			want: false,
		},
		{
			name: "test02",
			args: args{
				s: "(()",
			},
			want: false,
		},
		{
			name: "test03",
			args: args{
				s: "({})",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
