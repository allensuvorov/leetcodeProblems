func isPalindrome(s string) bool {
	l, r := 0, len(s)-1 // two pointers
	for l < r {         // rune

		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
		}

		// fmt.Print("left", string(s[l]), l)

		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			r--
		}

		// fmt.Println("right", string(s[r]), r)

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		l++
		r--

	}
	return true
}