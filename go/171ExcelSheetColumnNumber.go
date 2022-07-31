func titleToNumber(columnTitle string) int {
	pow := int(len(columnTitle) - 1)
	var sum int = 0
	for _, l := range columnTitle {
		sum += (int(l) - 'A' + 1) * int(math.Pow(26, float64(pow)))
		pow--
	}
	return int(sum)
}