func canPlaceFlowers(flowerbed []int, n int) bool {
    for i, v := range flowerbed {
        if v == 0 {
            leftIsEmpty := false
            rightIsEmpty := false
            if i == 0 || flowerbed[i - 1] == 0 {
                leftIsEmpty = true
            }
            if i == len(flowerbed) - 1 || flowerbed[i+1] == 0 {
                rightIsEmpty = true
            }
            if leftIsEmpty && rightIsEmpty {
                flowerbed[i] = 1
                n--
            }
        }
    }

    return n <= 0
}
