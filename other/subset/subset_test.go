package subset

import "testing"

type args struct {
	s1 []int
	s2 []int
}

func Test_isSubset(t *testing.T) {
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				s1: []int{1, 2, 3, 4, 5, 6},
				s2: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "test02",
			args: args{
				s1: []int{-1, -3, -2, 4, 1, 2, 3, 9},
				s2: []int{1, 9, 3},
			},
			want: true,
		},
		{
			name: "test03",
			args: args{
				s1: []int{-1, -3, 2, 1, 9},
				s2: []int{1, 2, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubset(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
