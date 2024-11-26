func findOrder(numCourses int, prerequisites [][]int) []int {
    ordering := make([]int, 0, numCourses)
    adjTable := make([][]int, numCourses)
    inds := make([]int, numCourses)

    for _, prereq := range prerequisites{
        a, b := prereq[0], prereq[1]
        adjTable[b] = append(adjTable[b], a)
        inds[a]++
    }

    q := []int{}

    for i, v := range inds {
        if v == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        now := q[0]
        q = q[1:]
        ordering = append(ordering, now)
        for _, nei := range adjTable[now] {
            inds[nei]--
            if inds[nei] == 0 {
                q = append(q, nei)
            }
        }
    }
    
    if len(ordering) == numCourses {
        return ordering
    }

    return nil
}
