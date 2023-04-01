// https://leetcode.com/problems/majority-element/
// Will implement here:
// Boyer-Moore Voting Algorithm
// Complexity Analysis

// Time complexity :
// O(n)

// Boyer-Moore performs constant work exactly n times, so the algorithm runs in linear time.

// Space complexity :

// O(1)

// Boyer-Moore allocates only constant additional memory.

func majorityElement(nums []int) int {
	var ctr int
	var mgr int
	for _, num := range nums {
		if ctr == 0 {
			mgr = num
		}
		if num == mgr {
			ctr++
		} else {
			ctr--
		}
	}
	return mgr
}