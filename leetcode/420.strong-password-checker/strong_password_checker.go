package leetcode

func strongPasswordChecker(password string) int {
	// missing character types
	lower, upper, digits, missingType := 1, 1, 1, 3
	for i := range password {
		if 0 < lower && 'a' <= password[i] && password[i] <= 'z' {
			lower--
			missingType--
		}
		if 0 < upper && 'A' <= password[i] && password[i] <= 'Z' {
			upper--
			missingType--
		}
		if 0 < digits && '0' <= password[i] && password[i] <= '9' {
			digits--
			missingType--
		}
		if missingType == 0 {
			break
		}
	}

	var replace, ones, twos int

	for p := 0; p+2 < len(password); p++ {
		if password[p] != password[p+1] || password[p+1] != password[p+2] {
			continue
		}

		// counting the length of consecutive repeats
		repeatingLen := 2
		for p+2 < len(password) && password[p] == password[p+2] {
			repeatingLen++
			p++
		}

		// strings that are repeated three times need to be replaced with another character
		replace += repeatingLen / 3
		if repeatingLen%3 == 0 {
			// if the length of the string exceeds 20,
			// can reduce the number of substitutions by removing 1 duplicate character
			ones++
		} else if repeatingLen%3 == 1 {
			// if the length of the string exceeds 20, you can reduce the number of substitutions
			// by removing 2 consecutive repetitive characters
			twos++
		}
	}

	// not enough long
	// use the missing type
	if len(password) < 6 {
		return max(missingType, 6-len(password))
	}

	// enough long
	// priority replacement for missing types
	if len(password) <= 20 {
		return max(missingType, replace)
	}

	// length > 20, deletion
	deletion := len(password) - 20
	replace -= min(deletion, ones)
	replace -= min(max(deletion-ones, 0), twos*2) / 2
	replace -= max(deletion-ones-2*twos, 0) / 3

	return deletion + max(missingType, replace)
}

func min(a, b int) int {
	return a + (b-a)>>31&(b-a)
}

func flip(a int) int {
	return a ^ 1
}

func sign(a int) int {
	// return flip(a >> 31 &1)
	a = a >> 31 & 1
	return a ^ 1
}

func max(a, b int) int {
	c := sign(a - b)
	d := flip(c)

	return a*c + b*d
}
