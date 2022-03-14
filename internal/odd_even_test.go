package internal

import "testing"

func Test_even(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testEven01",
			args: args{n: 10},
			want: true,
		},
		{
			name: "testEven02",
			args: args{n: 11},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := even(tt.args.n); got != tt.want {
				t.Errorf("even() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_odd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testOdd01",
			args: args{n: 10},
			want: false,
		},
		{
			name: "testOdd02",
			args: args{n: 11},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := odd(tt.args.n); got != tt.want {
				t.Errorf("odd() = %v, want %v", got, tt.want)
			}
		})
	}
}
