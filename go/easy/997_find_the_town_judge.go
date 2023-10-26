func findJudge(n int, trust [][]int) int {
    if n == 1 && len(trust) == 0 {
        return n
    }
    // map[int][]int - map[judge][]people
    trusteeMap := make(map[int][]bool)
    trustorSet := make(map[int]bool)

    for i := range trust {
        trustor := trust[i][0]
        trustee := trust[i][1]
        trustorSet[trustor] = true
        if _, ok := trusteeMap[trustee]; ok {
            trusteeMap[trustee][trustor] = true
        } else {
            people := make([]bool, n + 1)
            trusteeMap[trustee] = people
            trusteeMap[trustee][trustor] = true
        }
    } 

    for trustee, trustors := range trusteeMap {
        flag := true
        if trustorSet[trustee] { // judge not among trusters
            flag = false
            //fmt.Println("judge is a among trusters")
        }
        for i := 1; flag && i <= n; i++ { // go over all people
            if i != trustee && !trustors[i] { // if a person not trusting that judge
                flag = false
                //fmt.Println(i, " is not trust judge ", judge)
            }
        }
        if flag {
            return trustee
        }
    }
    return -1
}

//     judge
//    /  |  \
//   p1  p2  p3
