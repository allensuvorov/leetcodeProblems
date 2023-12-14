func finalPrices(prices []int) []int {
    res := make([]int, len(prices))
    var stack *node
    for i := range prices {
        for stack != nil && stack.Val >= prices[i] { // found discount
            res[stack.Ind] = stack.Val - prices[i]
            stack = stack.Next
        }
        newNode := node{prices[i], i, stack}
        stack = &newNode
    }
    for stack != nil {
        res[stack.Ind] = prices[stack.Ind]
        stack = stack.Next
    }
    return res
}

type node struct{
    Val int
    Ind int
    Next *node
}
