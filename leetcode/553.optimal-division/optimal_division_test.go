package leetcode

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testOptimalDivision01",
			args: args{nums: []int{1000, 100, 10, 2}},
			want: "1000/(100/10/2)",
		},
		{
			name: "testOptimalDivision02",
			args: args{nums: []int{3, 2}},
			want: "3/2",
		},
		{
			name: "testOptimalDivision03",
			args: args{nums: []int{3}},
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}
