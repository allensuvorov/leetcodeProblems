// https://leetcode.com/problems/majority-element/
// time: O(n)
// space: O(n)
func majorityElement(nums []int) int {
    hn := len(nums)/2
    m := make(map[int]int)
    // fill up the map
    for _, v := range nums{
        m[v]++ 
    }
    // find maj
    for k, v := range m {
        if v > hn {
            return k
        }
    }
    return 0
}

// Will implement here: Boyer-Moore Voting Algorithm
// Complexity Analysis
// Time complexity: O(n) // Boyer-Moore performs constant work exactly n times, so the algorithm runs in linear time.
// Space complexity: O(1) // Boyer-Moore allocates only constant additional memory.

func majorityElement(nums []int) int {
    count := 0
    major := 0
    for _, v := range nums {
        if count == 0 {
            major = v
        }
        if v == major {
            count++
        } else {
            count--
        }
    }
    return major
}
