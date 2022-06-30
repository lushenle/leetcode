package leetcode

import "testing"

func Test_numPrimeArrangements(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_numPrimeArrangements01",
			args: args{n: 5},
			want: 12,
		},
		{
			name: "test_numPrimeArrangements02",
			args: args{n: 6},
			want: 36,
		},
		{
			name: "test_numPrimeArrangements03",
			args: args{n: 7},
			want: 144,
		},
		{
			name: "test_numPrimeArrangements04",
			args: args{n: 100},
			want: 682289015,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPrimeArrangements(tt.args.n); got != tt.want {
				t.Errorf("numPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
