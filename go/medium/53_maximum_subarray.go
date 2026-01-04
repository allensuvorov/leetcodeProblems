// O(n) time
func maxSubArray(nums []int) int {
    curSum := 0
    curMax := nums[0]
    for _, v := range nums {
        curSum = max(v, v + curSum) // is that past helping?
        curMax = max(curMax, curSum)
    }
    return curMax
}

// O(n^2) time
func maxSubArray(nums []int) int {
    ans := math.MinInt
    for i := range nums {
        curSum := 0
        for j := i; j < len(nums); j++ {
            curSum += nums[j]
            ans = max(ans, curSum)
        }
    }
    return ans
}

// concurrent v1, synchronous channel - TLEs
func maxSubArray(nums []int) int {
    result := nums[0]
    ch := make(chan int)
    for i := range nums {
        go maxSum(nums[i:], ch)
    }
    for range nums {
        result = max(result, <- ch)
    }
    return result
}

func maxSum(nums []int, ch chan int) {
    result := nums[0]
    curSum := 0
    for _, v := range nums {
        curSum += v
        result = max(result, curSum)
    }
    ch <- result
}

// worker pool size NumCPU
import "runtime"
func maxSubArray(nums []int) int {
    n := len(nums)
    result := nums[0]

    input := make(chan []int, n)
    output := make(chan int, n)

    workerPoolSize := runtime.NumCPU()
    for range workerPoolSize {
        go worker(input, output)
    }

    for i := range nums {
        input <- nums[i:]
    }
    close(input)

    for range nums {
        result = max(result, <- output)
    }
    fmt.Printf("runtime.NumCPU(): %v \n", runtime.NumCPU())
    return result
}

func worker(input chan []int, output chan int) {
    for job := range input {
        output <- maxSum(job)
    }
}

func maxSum(nums []int) int {
    result := nums[0]
    curSum := 0
    for _, v := range nums {
        curSum += v
        result = max(result, curSum)
    }
    return result
}
