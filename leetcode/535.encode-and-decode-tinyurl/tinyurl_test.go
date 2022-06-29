package leetcode

import (
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	codec := Constructor()
	url := "https://shenle.lu/post/2022/02/some-exercises-and-golang-solutions/"
	encodeUrl := codec.encode(url)
	if !strings.EqualFold(encodeUrl, "http://tinurl.com/aHR0cHM6Ly9zaGVubGUubHUvcG9zdC8yMDIyLzAyL3NvbWUtZXhlcmNpc2VzLWFuZC1nb2xhbmctc29sdXRpb25zLw==") {
		t.Error("encode err")
	}

	decodeUrl := codec.decode(encodeUrl)
	if !strings.EqualFold(url, decodeUrl) {
		t.Error("decode err")
	}
}
