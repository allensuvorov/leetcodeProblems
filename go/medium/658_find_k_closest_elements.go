func findClosestElements(arr []int, k int, x int) []int {
    l := 0
    r := len(arr) - 1
    if x <= arr[l] {
        return arr[:k]
    }
    if x >= arr[r] {
        return arr[r-k+1:]
    }
    // case: arr[l] < x < arr[r]
    // find position of x with binary search
    ip := findInsertPosition(arr, x)
    lp := findLeftPosition(arr, ip, k, x)
    return arr[lp:lp + k]
}

func findInsertPosition (arr []int, x int) int {
    l := 0
    r := len(arr) - 1
    for l < r {
        m := l + (r - l)/2
        if arr[m] >= x {
            r = m
        } else {
            l = m + 1
        }
    }
    fmt.Println("InsertPosition:", l)
    return l
}

func findLeftPosition(arr []int, ip, k, x int) int {
    var i, l, r int
    if arr[ip] == x {
        l = ip - 1
        r = ip + 1
        i = 2
    } else {
        l = ip - 1
        r = ip
        i = 1
    }
    if k == 1 {
        closer := l
        if arr[r] - x < x - arr[l] {
            closer = r
        }
        return closer
    }
    for i = i; i <= k ; i++ {
        closer := l
        if arr[r] - x < x - arr[l] {
            closer = r
        } 
        switch closer {
        case l:
            if l == 0 {
                return l
            }
            l--
        case r:
            if r == len(arr)-1 {
                if r - k < 0 {
                    return 0
                }
                return r - k + 1
            }
            r++
        }
    }
    fmt.Println("leftPosition:", l)
    return l + 1
}
