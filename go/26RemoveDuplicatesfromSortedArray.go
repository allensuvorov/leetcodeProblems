package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
   // logic - find unique values and line them up
    var uniqsLen int = 0
    
    for i:=0; i<len(nums)-1; i++ {
        fmt.Println("i:", i, "uniqs:",uniqsLen)
        var first int = i
        for nums[first] == nums[i+1] {
            i++
            if i==len(nums)-1 {
                return uniqsLen+1
            }
        }
        uniqsLen++
        nums[uniqsLen] = nums[i+1]
    }
    return uniqsLen+1
}

func main() {
  fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
