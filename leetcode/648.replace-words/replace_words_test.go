package leetcode

import "testing"

func Test_replaceWords(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_replaceWords01",
			args{dictionary: []string{"cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery"},
			"the cat was rat by the bat",
		},
		{
			"test_replaceWords02",
			args{dictionary: []string{"a", "b", "c"}, sentence: "aadsfasf absbs bbab cadsfafs"},
			"a a b c",
		},
		{
			"test_replaceWords03",
			args{dictionary: []string{"a", "aa", "aaa", "aaaa"}, sentence: "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"},
			"a a a a a a a a bbb baba a",
		},
		{
			"test_replaceWords04",
			args{dictionary: []string{"catt", "cat", "bat", "rat"}, sentence: "the cattle was rattled by the battery"},
			"the cat was rat by the bat",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceWords(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("replaceWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
