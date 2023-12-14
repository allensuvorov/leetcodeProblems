type node struct{
    Val int
    Ind int
    Next *node
}

func finalPrices(prices []int) []int {
    var stack *node
    for i, discount := range prices {
        for stack != nil && stack.Val >= discount { // found discount
            prices[stack.Ind] -= discount
            stack = stack.Next
        }
        newNode := node{discount, i, stack}
        stack = &newNode
    }
    return prices
}
