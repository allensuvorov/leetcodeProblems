import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	// array, for loop, if

	fb := make([]string, n)

	for i := 0; i < n; i++ {
		fb[i] = ""
		if (i+1)%3 == 0 {
			fb[i] += "Fizz"
		}
		if (i+1)%5 == 0 {
			fb[i] += "Buzz"
		}
		if fb[i] == "" {
			fb[i] = strconv.Itoa(i)
			//fmt.Println(i)
		}
	}
	return fb
}

func main() {
	fmt.Println(fizzBuzz(10))
}
