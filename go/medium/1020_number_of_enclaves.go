func numEnclaves(grid [][]int) int {
	total := 0
	border := 0
	rows := len(grid)
	cols := len(grid[0])
	for r := range rows {
		for c := range cols {
			total += grid[r][c]
		}
	}

	fmt.Println("Before")
	for r := range rows {
		for c := range cols {
			fmt.Print(grid[r][c])
		}
		fmt.Println()
	}
	fmt.Println(total, border)

	for r := range rows {
		for c := range cols {
			if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
				if grid[r][c] == 1 {
					border += bfs(grid, r, c)
					if true {
						// fmt.Println("border > total")
						for r := range rows {
							for c := range cols {
								fmt.Print(grid[r][c])
							}
							fmt.Println()
						}
                        fmt.Println(total, border)
					}
				}
			}
		}
	}
	fmt.Println(total, border)
	return total - border
}

func bfs(grid [][]int, r, c int) int {
	count := 0
	q := [][]int{[]int{r, c}}
	for len(q) > 0 {
		size := len(q)
		count += size
		for range size {
			now := q[0]
			q = q[1:]
			r, c = now[0], now[1]
			grid[r][c] = 0
			if r > 0 && grid[r-1][c] == 1 {
				q = append(q, []int{r - 1, c})
			}
			if r < len(grid)-1 && grid[r+1][c] == 1 {
				q = append(q, []int{r + 1, c})
			}
			if c > 0 && grid[r][c-1] == 1 {
				q = append(q, []int{r, c - 1})
			}
			if c < len(grid[0])-1 && grid[r][c+1] == 1 {
				q = append(q, []int{r, c + 1})
			}
		}
	}
	return count
}
