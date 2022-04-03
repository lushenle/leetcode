package leetcode

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "testNextGreatestLetter01",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'a'},
			want: 'c',
		},
		{
			name: "testNextGreatestLetter02",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'c'},
			want: 'f',
		},
		{
			name: "testNextGreatestLetter03",
			args: args{letters: []byte{'c', 'f', 'j'}, target: 'd'},
			want: 'f',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
