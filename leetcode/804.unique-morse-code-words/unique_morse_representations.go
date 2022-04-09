package leetcode

func uniqueMorseRepresentations(words []string) int {
	m := map[string]string{"a": ".-", "b": "-...", "c": "-.-.",
		"d": "-..", "e": ".", "f": "..-.", "g": "--.", "h": "....",
		"i": "..", "j": ".---", "k": "-.-", "l": ".-..", "m": "--",
		"n": "-.", "o": "---", "p": ".--.", "q": "--.-", "r": ".-.",
		"s": "...", "t": "-", "u": "..-", "v": "...-", "w": ".--",
		"x": "-..-", "y": "-.--", "z": "--.."}

	ans := make(map[string]string)
	for _, word := range words {
		var morse string
		for _, v := range word {
			morse += m[string(v)]
		}

		ans[morse] = word
	}

	return len(ans)
}
