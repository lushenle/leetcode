package leetcode

import (
	"encoding/base64"
	"strings"
)

type Codec struct {
	base string
}

func Constructor() Codec {
	return Codec{base: "http://tinurl.com/"}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	str := c.base
	str += base64.StdEncoding.EncodeToString([]byte(longUrl))

	return str
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	url := strings.Split(shortUrl, c.base)
	data, err := base64.StdEncoding.DecodeString(url[1])
	if err != nil {
		return ""
	}

	return string(data)
}
