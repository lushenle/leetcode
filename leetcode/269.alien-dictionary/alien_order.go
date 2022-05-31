package leetcode

func alienOrder(words []string) string {
	graph, set, d := map[byte][]byte{}, map[byte]bool{}, make([]int, 26)

out:
	for i, word := range words {
		for j := range word {
			set[word[j]] = true
		}

		if i < len(words)-1 {
			for j := 0; j < min(len(words[i]), len(words[i+1])); j++ {
				if words[i][j] != words[i+1][j] {
					graph[words[i][j]] = append(graph[words[i][j]], words[i+1][j])
					d[words[i+1][j]-'a']++
					continue out
				}
			}
			if len(words[i]) > len(words[i+1]) {
				return ""
			}
		}
	}

	var ans []byte
	for s := range set {
		if d[s-'a'] == 0 {
			ans = append(ans, s)
		}
	}

	for i := 0; i < len(ans); i++ {
		for _, nxt := range graph[ans[i]] {
			d[nxt-'a']--
			if d[nxt-'a'] == 0 {
				ans = append(ans, nxt)
			}
		}
	}

	if len(ans) == len(set) {
		return string(ans)
	}

	return ""
}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}
