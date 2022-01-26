package leetcode

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{3, 2, 2, 3}, val: 3},
			want: 2,
		},
		{
			name: "test02",
			args: args{nums: []int{1, 1, 1, 1, 1, 1}, val: 1},
			want: 0,
		},
		{
			name: "test03",
			args: args{nums: []int{1, 2, 3, 4, 5}, val: 1},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
