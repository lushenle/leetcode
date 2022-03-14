package internal

func odd(n int) bool {
	return n&1 == 1

}

func even(n int) bool {
	return n&1 == 0
}
