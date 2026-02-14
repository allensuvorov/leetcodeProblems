func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	g := make(map[int][]int) // adj table
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	bobTime := make(map[int]int) // node to time
	seen := make(map[int]bool)
	const startTime = 0
	const alice = 0
	res := math.MinInt
	
    bobDFS(g, bobTime, seen, bob, startTime)
	seen = make(map[int]bool)
	aliceDFS(g, amount, bobTime, seen, alice, startTime, 0, &res)
	
    return res
}

func bobDFS(
	g map[int][]int,
	bobTime map[int]int,
	seen map[int]bool,
	node, time int,
) bool {
	res := node == 0
	seen[node] = true
	for _, nei := range g[node] {
		if !seen[nei] && bobDFS(g, bobTime, seen, nei, time+1) {
			res = true
		}
	}
	if res {
		bobTime[node] = time
	}
	return res
}

func aliceDFS(
	g map[int][]int,
	amount []int,
	bobTime map[int]int,
	seen map[int]bool,
	node, time, curProfit int,
	maxProfit *int,
) {
	seen[node] = true
	if bTime, ok := bobTime[node]; !ok || bTime > time {
		curProfit += amount[node]
	}

	if bobTime[node] == time {
		curProfit += amount[node] / 2
	}

	if node != 0 && len(g[node]) == 1 { // leaf
		*maxProfit = max(*maxProfit, curProfit)
	}
	for _, nei := range g[node] {
		if !seen[nei] {
			aliceDFS(g, amount, bobTime, seen, nei, time+1, curProfit, maxProfit)
		}
	}
}
