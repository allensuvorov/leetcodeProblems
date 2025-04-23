func kWeakestRows(mat [][]int, k int) []int {
    n := len(mat)
    countSoldiers := func(nums []int) int {
        l, r := 0, len(nums) - 1
        for l <= r {
            m := l + (r - l) / 2
            if nums[m] == 1 {
                l = m + 1
            } else {
                r = m - 1
            }
        }
        return r + 1
    }
    power := make([][]int, n)
    for i, row := range mat {
        power[i] = []int{countSoldiers(row), i}
    }

    slices.SortFunc(power, func(a, b []int)int {
        if a[0] == b[0] {
            return a[1] - b[1]
        }
        return a[0] - b[0]
    })
    weakestRows := make([]int, k)
    for i := range k {
        weakestRows[i] = power[i][1]

    }
    return weakestRows
}
