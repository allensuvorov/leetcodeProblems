func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerCount := 0
	for i := range flowerbed {
		curr := flowerbed[i] == 0
		prev := i == 0 || flowerbed[i-1] == 0
		next := i == len(flowerbed)-1 || flowerbed[i+1] == 0
		if prev && curr && next {
			flowerbed[i] = 1
			flowerCount++
		}
	}
	return flowerCount >= n
}
