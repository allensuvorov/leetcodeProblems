func addBinary(a string, b string) string {
	var carry rune = '0'
	var sum rune
	var res []rune
	var i int
	if len(b) > len(a) {
		a, b = b, a
	}
	var dif int = len(a) - len(b)

	for i = len(a) - 1; i > dif-1; i-- {
		//fmt.Println("loop1 start", "i=", i, "dif=", dif)
		if carry == '0' {
			if a[i] == '1' && b[i-dif] == '1' {
				sum, carry = '0', '1'
			}
			if a[i] == '1' && b[i-dif] == '0' {
				sum, carry = '1', '0'
			}
			if a[i] == '0' && b[i-dif] == '1' {
				sum, carry = '1', '0'
			}
			if a[i] == '0' && b[i-dif] == '0' {
				sum, carry = '0', '0'
			}
		} else { // carry = '1'
			if a[i] == '1' && b[i-dif] == '1' {
				sum, carry = '1', '1'
			}
			if a[i] == '1' && b[i-dif] == '0' {
				sum, carry = '0', '1'
			}
			if a[i] == '0' && b[i-dif] == '1' {
				sum, carry = '0', '1'
			}
			if a[i] == '0' && b[i-dif] == '0' {
				sum, carry = '1', '0'
			}
		}

		res = append([]rune{sum}, res...)
		//fmt.Println("Loop1, a+b, sum = ", string(sum), "carry=", carry, "i=", i,)
	}
	fmt.Println("a+b loop res =", string(res), "carry=", carry, "i=", i, "dif=", dif)

	for i = dif - 1; i >= 0 && carry == '1'; i-- {
		if a[i] == '1' {
			sum, carry = '0', '1'
		}

		if a[i] == '0' {
			sum, carry = '1', '0'
		}
		res = append([]rune{sum}, res...)

	}
	fmt.Println("a+carry loop res =", string(res), "carry=", carry)

	if carry == '1' {
		res = append([]rune{carry}, res...)
	}

	if i >= 0 {
		res = append([]rune(a[:i+1]), res...)
	}
	fmt.Println("final res =", string(res), "carry=", carry)
	return string(res)
}