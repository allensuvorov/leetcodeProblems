// 4 ms
func abs(n int) int{
    if n < 0 {
        return -n
    }
    return n
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
    ans := 0
    for i := 0; i < len(arr) - 2; i++ {
        for j:= i + 1; j < len(arr) - 1; j++ {
            if abs(arr[i] - arr[j]) <= a {
                for k := j + 1; k < len(arr); k++ {
                    if abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c{
                        ans++
                    }
                }
            }
        }
    }
    return ans
}

// 500 ms

func countGoodTriplets(arr []int, a int, b int, c int) int {
    ans := 0
    i, j, k := 0, 1, 2
    for i < len(arr) - 2 {
        if j == len(arr) - 1 {
            i++
            j = i + 1
            k = j + 1
            continue
        }

        if k == len(arr) {
            j++
            k = j + 1
            continue
        }

        if abs(arr[i] - arr[j]) > a {
            j++
            k = j + 1
            continue
        }

        if abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k]) <= c {
            fmt.Println(arr[i], arr[j], arr[k])
            ans++
        } 
        k++
    }
    return ans
}

func abs(n int) int{
    if n < 0 {
        return -n
    }
    return n
}
