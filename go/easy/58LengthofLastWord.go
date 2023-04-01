func lengthOfLastWord(s string) int {
	var res, i int

	// option 1
	// for i = len(s)-1; i >= 0; i-- {
	//     if s[i] != ' ' {
	//         res++
	//     } else if res != 0 {
	//         break
	//     }
	// }

	// option 2
	for i = len(s) - 1; s[i] == ' '; i-- {
	}

	for ; s[i] != ' '; i-- {
		res++
		if i == 0 {
			break
		}
	}

	return res
}