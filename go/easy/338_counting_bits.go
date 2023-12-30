my solution - has human logic
func countBits(n int) []int {
    // 0, 1,    2, 3,    4, 5, 6, 7,    8, 9, 10, 11, 12, 13, 14, 15,   16 
    // 0, 1,    1, 2,    1, 2, 2, 3,    1, 2,  2,  3,  2,  3,  3,  4,   1

    ans := make([]int, n + 1)
    maxPowOfTwoProduct := 0
    for i := 1; i <= n; i++ {        
        // calculate first values: 0, 1, 1, 2
        if i == 1 {
            maxPowOfTwoProduct = 1
            ans[i] = 1
            continue
        }
        if i == 2 {
            maxPowOfTwoProduct = 2
            ans[i] = 1
            continue
        }
        
        // calculate following values:
        if i == maxPowOfTwoProduct * 2 {
            maxPowOfTwoProduct *= 2
            ans[i] = 1
        } else {
            ans[i] = ans[maxPowOfTwoProduct] + ans[i - maxPowOfTwoProduct]
        }
    }

    return ans
}

// short, but not very intuitive 
func countBits(n int) []int {
     bitArr := make([]int, n+1)
     bitArr[0] = 0
	   for i := 1; i <= n; i++ {
		     if i%2 == 0 {
			       bitArr[i] = bitArr[i/2]
		     } else {
			       bitArr[i] = bitArr[i-1] + 1
		     }
	   }
	   return bitArr
}
