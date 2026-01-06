func singleNumber(nums []int) int {
    ans := 0
    for _, v := range nums {
        ans ^= v
    }
    return ans
}

// concurrency with channels
// [][][] -> workers -> output chan -> consumer
func singleNumber(nums []int) int {
    n := len(nums)
    workerCount := runtime.NumCPU()
    jobSize := max(1, n / workerCount)
    output := make(chan int, workerCount)

    for i := 0; i < n; i += jobSize {
        job := nums[i: min(n, i+jobSize)]
        go singleNumberWorker(job, output)    
    }

    single := 0
    for i := 0; i < n; i += jobSize {
        single = single ^ <-output
    }
    return single
}

func singleNumberWorker(nums []int, output chan<- int) {
    single := 0
    for _, v := range nums {
        single = single ^ v
    }
    output <- single
}
