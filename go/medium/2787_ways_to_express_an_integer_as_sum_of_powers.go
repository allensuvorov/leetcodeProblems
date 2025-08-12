func numberOfWays(n int, x int) int {
    terms := []int{}
    for i := 1; i <= n; i++ {
        pow := i
        for range x - 1 {
            pow = pow * i
        }
        if pow <= n {
            terms = append(terms, pow)
        } else {
            break
        }
    }
    fmt.Println(terms)
    return pathSum(terms, n) // dfs
}


func pathSum(terms []int, sum int) int {
    if sum == 0 {
        return 1
    }

    if sum < 0 {
        return 0
    }

    if len(terms) == 0 {
        return 0
    }
    count := 0
    for i, v := range terms {
        count += pathSum(terms[i+1:], sum - v)
    }
    return count
}
