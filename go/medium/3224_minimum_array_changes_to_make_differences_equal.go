func minChanges(nums []int, k int) int {
    if k == 0 {
        return 0
    }
    ans := math.MaxInt
    diffFrequencyCount := make(map[int]int)
    singleChangeCapacityFrequencyCount := make([]int, k + 1)

    for l, r := 0, len(nums) - 1; l < r; l, r = l+1, r-1 {
        absDiff := absDiff(nums[l], nums[r])
        diffFrequencyCount[absDiff]++
        singleChangeCapacity := max(nums[l], nums[r], k - nums[l], k - nums[r])
        singleChangeCapacityFrequencyCount[singleChangeCapacity]++
    }
    
    prefixSum := make([]int, k + 1)
    for sum, i := 0, k; i >= 0; i-- {
        v := singleChangeCapacityFrequencyCount[i]
        sum += v
        prefixSum[i] = sum
    }

    for diff, freqCount := range diffFrequencyCount {
        singleChangeCount := (prefixSum[diff] - freqCount)
        doubleChangeCount := len(nums)/2 - singleChangeCount - freqCount
        totalChangeCount := singleChangeCount + doubleChangeCount * 2
        ans = min(ans, totalChangeCount)
    }
    return ans

}

func absDiff(a, b int) int {
    c := a - b
    if c < 0 {
        c = -c
    }
    return c
}
/*
The main idea/subproblem is to compute in constant time the number of changes needed for a given X using extra memory for a prefix sum array.

First. Let's count differences that actually exist. Those show how many pairs for a given X need ZERO change.

Second. Looking at a pair and the boundaries, can we tell the maximum X that we can squeeze after a SINGLE change?
- Yes, here is an example for a pair [1, 2] with boundaries [0,4]:

4 <- upper boundary (K)
3
2 <- num2
1 <- num1
0 <- lower boundary

Here, if we change num2, the max distance we can get till we hit a boundary is 3: (4 - num1) = 3. So the "capacity" of this pair after a single change is 3.

Now, pairs with capacity 3 and higher all can handle X = 3. Hence we calculate prefixSum array, which shows the total of how many pairs can handle X after a single change. (One other subproblem for me was to realise that a prefix sum includes zeroChangeCount)

Finally, we can do the math. Here is what we know:
- the total number of pairs  = n/2.
- zeroChangeCount, which in code is named as diffFreqencyCount
- singleChangeCount, which is prefixSum minus the zeroChangeCount for a each existing X
- doubleChangeCount is the rest of the pairs.

Thus, for each X: totalChangeCount := singleChangeCount + doubleChangeCount * 2

Let me know if you have any questions. Hope that helps. Will be happy to do a walk through the code.
*/
