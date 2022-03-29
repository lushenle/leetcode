package leetcode

import "testing"

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "Test_maxConsecutiveAnswers01",
			args:    args{answerKey: "TTFF", k: 2},
			wantAns: 4,
		},
		{
			name:    "Test_maxConsecutiveAnswers02",
			args:    args{answerKey: "TFFT", k: 1},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
