package internal

import (
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "test02",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.nums...); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "test02",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxElement(tt.args.nums); got != tt.want {
				t.Errorf("maxElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "test02",
			args: args{nums: []int{-1, -2, -3, -4, -5}},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.nums...); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaxElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxElement(generate())
	}
}

func BenchmarkMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		max(1, 2, 3, 4)
	}
}

func BenchmarkMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		min(1, 2, 3, 4)
	}
}
