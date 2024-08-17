type NeighborSum struct {
    Grid [][]int
}


func Constructor(grid [][]int) NeighborSum {
    return NeighborSum{grid}
}


func (this *NeighborSum) Position(value int) []int{
    for i := range this.Grid {
        for j := range this.Grid {
            if this.Grid[i][j] == value {
                return []int{i, j}
            }
        }
    }
    return nil
}

func (this *NeighborSum) AdjacentSum(value int) int {
    sum := 0
    pos := this.Position(value)
    if pos[0] > 0 {
        sum += this.Grid[pos[0] - 1][pos[1]]
    }
    if pos[0] < len(this.Grid) - 1 {
        sum += this.Grid[pos[0] + 1][pos[1]]
    }
    if pos[1] > 0 {
        sum += this.Grid[pos[0]][pos[1] - 1]
    }
    if pos[1] < len(this.Grid) - 1 {
        sum += this.Grid[pos[0]][ pos[1] + 1]
    }
    return sum
}


func (this *NeighborSum) DiagonalSum(value int) int {
    sum := 0
    pos := this.Position(value)
    if pos[0] > 0 && pos[1] > 0 {
        sum += this.Grid[pos[0] - 1][pos[1] - 1]
    }
    if pos[0] < len(this.Grid) - 1 && pos[1] < len(this.Grid) - 1 {
        sum += this.Grid[pos[0] + 1][pos[1] + 1]
    }
    if pos[0] > 0 && pos[1] < len(this.Grid) - 1 {
        sum += this.Grid[pos[0] - 1][pos[1] + 1]
    }
    if pos[0] < len(this.Grid) - 1 && pos[1] > 0 {
        sum += this.Grid[pos[0] + 1][pos[1] - 1]
    }
    return sum
}


/**
 * Your NeighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
