// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func reverse(x int) int {

	var (
		num	int = x // max 2147483647 min -2147483648
		rev   	int = 0
		pop 	int
	)
	const (
		maxInt int = int(^uint32(0) >> 1)
		minInt int = -(maxInt - 1)
	)
	//fmt.Println(maxInt)

	if num < 0 {
		num = -num
	}

	for num != 0 {
		pop = num % 10 // get end digit
		
		if rev > maxInt/10 or (rev == maxInt/10 and pop > 7) {
			return 0
			}

		# check if rev is about to overflow
            if rev > (int_max/10) or (rev == int_max/10 and pop > 7):
                return 0
            if rev < (int_min/10) or (rev == int_min/10 and pop < -8):
                return 0



		num = int(num / 10) // pop
	}
	fmt.Println(endDigit)
	return revNum
}

func main() {
	fmt.Println(reverse(10))
}
