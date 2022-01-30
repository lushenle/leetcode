package leetcode

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test01",
			args: args{s1: "this apple is sweet", s2: "this apple is sour"},
			want: []string{"sweet", "sour"},
		},
		{
			name: "test02",
			args: args{s1: "i am trump", s2: "i am biden"},
			want: []string{"trump", "biden"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
