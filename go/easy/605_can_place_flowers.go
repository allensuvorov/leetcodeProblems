func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := range flowerbed {
        if flowerbed[i] == 0 { // current
            if i == 0 || flowerbed[i - 1] == 0 { // prev
                if i == len(flowerbed) - 1 || flowerbed[i+1] == 0 { // next
                    flowerbed[i] = 1
                    n--
                }
            }
        }
    }

    return n <= 0
}
