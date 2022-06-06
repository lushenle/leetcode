package leetcode

import "strings"

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, email := range emails {
		sp := strings.Split(email, "@")
		localName := strings.ReplaceAll(strings.Split(sp[0], "+")[0], ".", "")
		m[localName+"@"+sp[1]] = true
	}

	return len(m)
}
