func convertToTitle(columnNumber int) string {
	n := columnNumber
	var s []rune
	var digit int

	for n > 0 {
		if n%26 != 0 {
			digit = n % 26 // get remainder
			n = n - digit
		} else {
			digit = 26 // append Z
			n = n - 26
		}
		s = append([]rune{rune(digit + 64)}, s...)
		n = n / 26
	}

	fmt.Println(string(s), n, digit)
	return string(s)
}

// (D2*26^2 + D1*26^1)/26 = (D2*26^1 + D1)