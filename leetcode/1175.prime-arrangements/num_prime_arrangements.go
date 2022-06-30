package leetcode

func numPrimeArrangements(n int) int {
	cntPrime, cntComposite, ans := 0, 1, 1
	const MOD int = 1e9 + 7

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			cntPrime++
			ans = ans * cntPrime % MOD
		} else {
			cntComposite++
			ans = ans * cntComposite % MOD
		}
	}

	return ans % MOD
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
