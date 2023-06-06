func nextGreaterElement(nums1 []int, nums2 []int) []int {
    ans := make([]int, len(nums1))
    m := make(map[int]int)

    for i := range nums2 {
        m[nums2[i]] = i
    }

    for i, v:= range nums1 {    
        ans[i] = findNext(nums2, m[v], v)
    }
    return ans
}

func findNext(arr []int, i, n int) int {
    for j := i; j < len(arr); j ++ {
        if arr[j] > n {
            return arr[j]
        }
    }
    return - 1
}


// O(n+m)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    var nge = func(nums []int) map[int]int {
        ngeMap := make(map[int]int)
        var stack []int
        
        for _, num := range nums {
            for len(stack) > 0 && stack[len(stack) - 1] < num {
                top := stack[len(stack) - 1]
                ngeMap[top] = num
                stack = stack[:len(stack) - 1]
            }
            stack = append(stack, num)
        }

        for _, num := range stack {
            ngeMap[num] = -1
        }

        return ngeMap
    }

    ngeMap := nge(nums2)

    var result []int
    for _, num := range nums1 {
        result = append(result, ngeMap[num])
    }

    return result
}
