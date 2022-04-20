package leetcode

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testToGoatLatin01",
			args: args{sentence: "I speak Goat Latin"},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name: "testToGoatLatin02",
			args: args{sentence: "over"},
			want: "overmaa",
		},
		{
			name: "testToGoatLatin03",
			args: args{sentence: "The quick brown fox jumped over the lazy dog"},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.sentence); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
