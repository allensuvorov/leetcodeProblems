func numOfUnplacedFruits(fruits []int, baskets []int) int {
    count := 0

    for _, fruitCount := range fruits {
        allocated := false
        for j, basketSize := range baskets {
            if basketSize >= fruitCount {
                baskets[j] = 0
                allocated = true
                break
            }
        }

        if !allocated {
            count++
        }
    }
    return count
}
