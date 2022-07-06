package leetcode

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_evaluate01",
			args: args{expression: "(let x 2 (mult x (let x 3 y 4 (add x y))))"},
			want: 14,
		},
		{
			name: "test_evaluate02",
			args: args{expression: "(let x 3 x 2 x)"},
			want: 2,
		},
		{
			name: "test_evaluate03",
			args: args{expression: "(let x 1 y 2 x (add x y) (add x y))"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.expression); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
