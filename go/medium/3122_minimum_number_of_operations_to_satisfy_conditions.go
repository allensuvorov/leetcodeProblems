func minimumOperations(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    prev := [10]int{}
    curr := [10]int{}

    for c := range n {
        for r := range m { // curr count 
            curr[grid[r][c]]++
        }
        for i := range 10 { // curr count + what was best in prev for any other number
            maxPrev := 0
            for j := range 10 {
                if i != j {
                    maxPrev = max(maxPrev, prev[j])
                }
            }
            curr[i] += maxPrev
        }
        prev = curr
        curr = [10]int{}
    }

    maxPrev := 0
    for i := range 10 {
        maxPrev = max(maxPrev, prev[i])
    }
    fmt.Println(prev)
    return m * n - maxPrev
}


/*

1 0 2
1 0 2

0 [0, 4, 4]
1 [2, 0, 4]
2 [0, 2, 6]

*/
