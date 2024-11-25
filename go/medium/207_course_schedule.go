func canFinish(numCourses int, prerequisites [][]int) bool {
    adjTable := make([][]int, numCourses)
    indegree := make([]int, numCourses)
    for _, prereq := range prerequisites{
        a, b := prereq[0], prereq[1]
        adjTable[b] = append(adjTable[b], a)
        indegree[a]++
    }

    canFinishCount := 0
    q := []int{}
    for i, v := range indegree {
        if v == 0 {
            canFinishCount++
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        now := q[0]
        q = q[1:]
        for _, nei := range adjTable[now]{
            indegree[nei]--
            if indegree[nei] == 0 {
                canFinishCount++
                q = append(q, nei)
            }
        }
    }

    return canFinishCount == numCourses
}
