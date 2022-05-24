package leetcode

type AuthenticationManager struct {
	tokens     map[string]int
	timeToLive int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		tokens:     map[string]int{},
		timeToLive: timeToLive,
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	am.tokens[tokenId] = currentTime
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := am.tokens[tokenId]; ok {
		if v+am.timeToLive > currentTime {
			am.tokens[tokenId] = currentTime
		}
	}
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	ans := 0
	for _, v := range am.tokens {
		if v+am.timeToLive > currentTime {
			ans++
		}
	}

	return ans
}
