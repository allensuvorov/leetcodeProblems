
type NestedIterator struct {
    currentIndex int
    q []*NestedInteger
    flatList []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    // flatten with iterative DFS with a q
    nums := []int{}
    // parseNested recursive
    var dfs func(node *NestedInteger) 
    dfs = func(node *NestedInteger) {
        if node == nil {
            return
        }
        if node.IsInteger() { // it is a leaf
            nums = append(nums, node.GetInteger())
        } else { // it is a parent
            for _, v := range node.GetList() {
                dfs(v)
            }
        }
    }

    for _, node := range nestedList {
        dfs(node)
    }
    return &NestedIterator{
        currentIndex: -1,
        flatList: nums,
    }
}

func (this *NestedIterator) Next() int {
    this.currentIndex++
    return this.flatList[this.currentIndex]
}

func (this *NestedIterator) HasNext() bool {
    return len(this.flatList) > this.currentIndex + 1
}
