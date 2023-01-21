func canPlaceFlowers(flowerbed []int, n int) bool {
    count := 0
    for i := range flowerbed{
        if flowerbed[i] == 0 {
            if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0 ) {
                count++
                flowerbed[i] = 1
            } 
        }
    }
    return count >= n
}
