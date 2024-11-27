/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    n := len(grid)
    val := true
    first := grid[0][0]
    if first == 0 {
        val = false
    }
    isLeaf := true
    for r := range n {
        for c := range n {
            if grid[r][c] != first {
                isLeaf = false
            }
        }
    }
    node := &Node{
        Val: val,
        IsLeaf: isLeaf,
    }
    if isLeaf {
        return node
    }

    node.TopLeft = construct(subGrid(grid, 0, n/2, 0, n/2))
    node.TopRight = construct(subGrid(grid, 0, n/2, n/2, n))
    node.BottomLeft = construct(subGrid(grid, n/2, n, 0, n/2))
    node.BottomRight = construct(subGrid(grid, n/2, n, n/2, n))

    return node
}

func subGrid(grid [][]int, rowStart, rowEnd, colStart, colEnd int) [][]int{
    subGrid := make([][]int, 0, len(grid) / 2)
    for r := rowStart; r < rowEnd; r++ {
        subGrid = append(subGrid, grid[r][colStart:colEnd])
    }
    return subGrid
}
