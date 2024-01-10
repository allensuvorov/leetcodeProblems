func minMutation(startGene string, endGene string, bank []string) int {
    // if startGene == endGene {
    //     return 0
    // }
    visited := make([]bool, len(bank))
    q := []mut{{startGene, 0}}

    for len(q) > 0 {
        for i, bankGene := range bank {
            if !visited[i] && bankGene != startGene {
                if isOneMut(q[0].gene, bankGene) {
                    if bankGene == endGene {
                        return q[0].n + 1
                    }
                    visited[i] = true
                    q = append(q, mut{bankGene, q[0].n + 1})
                }
            }
        }
        q = q[1:]
    }
    return -1
}

type mut struct {
    gene string
    n int
}

func isOneMut(aGene, bGene string) bool {
    count := 0
    for i := range aGene {
        if aGene[i] != bGene[i] {
            count++
        }
    }
    return count == 1
}
