func sumOddLengthSubarrays(arr []int) int {
    // n^2 with sliding window
    totalSum := 0
    for window := 1; window <= len(arr); window += 2 {
        windowSum := 0
        for i := range arr {
            windowSum += arr[i]
            if i + 1 > window { // cut tailing value
                windowSum -= arr[i - window]
            }
            if i + 1 >= window {
                totalSum += windowSum
            } 
        }
    }
    return totalSum
}

// concurrent with a mutex
func sumOddLengthSubarrays(arr []int) int {
    // n^2 with sliding window
    // [workers] -> var with a mutex
    totalSum := 0
    var mu sync.Mutex
    var wg sync.WaitGroup
    for window := 1; window <= len(arr); window += 2 {
        wg.Add(1)
        go func() {
            sumSubarrays(arr, window, &mu, &totalSum)
            wg.Done()
        }()
    }

    wg.Wait()
    return totalSum
}

func sumSubarrays(arr []int, window int, mu *sync.Mutex, totalSum *int) {
    res := 0
    windowSum := 0
    for i := range arr {
        windowSum += arr[i]
        if i + 1 > window { // cut tailing value
            windowSum -= arr[i - window]
        }
        if i + 1 >= window {
            res += windowSum
        } 
    }
    mu.Lock()
    *totalSum += res
    mu.Unlock()
}

// concurrent with a buffered channel
func sumOddLengthSubarrays(arr []int) int {
    // n^2 with sliding window
    // [workers] -> ch -> [consumer]
    n := len(arr)
    totalSum := 0
    ch := make(chan int, n/2)
    var wg sync.WaitGroup
    for window := 1; window <= len(arr); window += 2 {
        wg.Add(1)
        go func() {
            sumSubarrays(arr, window, ch)
            wg.Done()
        }()
    }
    go func() {
        wg.Wait()
        close(ch)
    }()
    for subSum := range ch {
        totalSum += subSum
    }
    return totalSum
}

func sumSubarrays(arr []int, window int, ch chan int) {
    res := 0
    windowSum := 0
    for i := range arr {
        windowSum += arr[i]
        if i + 1 > window { // cut tailing value
            windowSum -= arr[i - window]
        }
        if i + 1 >= window {
            res += windowSum
        } 
    }
    ch <- res
}
