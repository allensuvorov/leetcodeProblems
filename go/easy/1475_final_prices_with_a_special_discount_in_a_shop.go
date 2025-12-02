// stack via linked list
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

// stack - via slice
func finalPrices(prices []int) []int {
    stack := []int{}
    discPrices := make([]int, len(prices))
    for i, v := range prices {
        discPrices[i] = v
        for len(stack) > 0 && prices[stack[len(stack)-1]] >= v {
            discPrices[stack[len(stack)-1]] -= v 
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return discPrices
}
