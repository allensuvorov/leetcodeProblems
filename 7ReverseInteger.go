// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func reverse(x int) int {

	var (
		num int = x
		rev int = 0
		pop int
	)
	const (
		maxInt int = int(^uint32(0) >> 1) // max 2147483647 
		minInt int = -(maxInt - 1) // min -2147483648
	)
	//fmt.Println(maxInt)

	if x < 0 {
		num = -num
	}

	for num != 0 {
		pop = num % 10 // get end digit
		fmt.Println(pop)

		// check if rev is about to overflow
		if rev > maxInt/10 || (rev == maxInt/10 && pop > 7) || rev < minInt/10 ||
			(rev == minInt/10 && pop < -8) {
			return 0
		}

		// push
		rev = rev*10 + pop

		// pop
		num = int(num / 10)
	}
	if x < 0 {
		rev = -rev
	}

	return rev
}

func main() {
	fmt.Println(reverse(-123))
}
