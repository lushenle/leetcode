package leetcode

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name:    "testSelfDividingNumbers01",
			args:    args{left: 1, right: 22},
			wantAns: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
		{
			name:    "testSelfDividingNumbers02",
			args:    args{left: 47, right: 85},
			wantAns: []int{48, 55, 66, 77},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := selfDividingNumbers(tt.args.left, tt.args.right); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("selfDividingNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}

}
