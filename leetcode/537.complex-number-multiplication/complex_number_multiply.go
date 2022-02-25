package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	realA, imagA := parse(num1)
	realB, imagB := parse(num2)
	r := realA*realB - imagA*imagB
	i := realA*imagB + realB*imagA
	return strconv.Itoa(r) + "+" + strconv.Itoa(i) + "i"
}

func parse(s string) (int, int) {
	ss := strings.Split(s, "+")
	r, _ := strconv.Atoi(ss[0])
	i, _ := strconv.Atoi(ss[1][:len(ss[1])-1])
	return r, i
}

func complexNumberMultiply1(num1, num2 string) string {
	x, _ := strconv.ParseComplex(num1, 64)
	y, _ := strconv.ParseComplex(num2, 64)
	return fmt.Sprintf("%d+%di", int(real(x*y)), int(imag(x*y)))
}
