func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerCount := 0
	prev, next := 0, 0
	for i := range flowerbed {
		if i > 0 {
			prev = flowerbed[i-1]
		}
		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}
		if prev+flowerbed[i]+next == 0 {
			flowerbed[i] = 1
			flowerCount++
		}
	}
	return flowerCount >= n
}
