func distMoney(money int, children int) int {
    if money < children {
        return -1
    }

    if money % children == 0 && money / children == 8 {
        return children
    }

    ans := 0
    for children > 0 && money - 8 != 4 && money - children >= 8  {
        money -= 8
        children--
        ans++
    }
    return ans
}

// money 16, children = 3
// 12 : 2 - 0

// 18 : 10 + 8

// 8 : 1
// 8 + 1 : 2
