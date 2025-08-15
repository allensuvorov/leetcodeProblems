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

// early exit
func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := range flowerbed {
        if flowerbed[i] != 0 { // current
            continue
        }

        if i != 0 && flowerbed[i - 1] == 1 { // prev
            continue
        }

        if i != len(flowerbed) - 1 && flowerbed[i+1] == 1 { // next
            continue
        }

        flowerbed[i] = 1
        n--
    }

    return n <= 0
}
