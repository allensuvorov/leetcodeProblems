// someone's solution that's 0ms

func init() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, math.MaxInt32)
	_out, _ := os.Create("user.out")
	out := bufio.NewWriter(_out)
	for in.Scan() {
		s := in.Bytes()
		n := bytes.Count(s, []byte{','}) + 1
		ans, l, r, i, j, readL, readR := 0, 0, n-1, 0, len(s)-1, true, true
		var x, y int
		for l < r {
			if readL {
				readL = false
				i++
				x = int(s[i] & 15)
				i++
				for s[i]&15 < 10 {
					x = x*10 + int(s[i]&15)
					i++
				}
			}
			if readR {
				readR = false
				for j -= 2; s[j]&15 < 10; j-- {
				}
				k := j + 1
				y = int(s[k] & 15)
				k++
				for s[k]&15 < 10 {
					y = y*10 + int(s[k]&15)
					k++
				}
			}
			ans = max(ans, min(x, y)*(r-l))
			if x < y {
				l++
				readL = true
			} else {
				r--
				readR = true
			}
		}
		fmt.Fprintln(out, ans)
	}
	out.Flush()
	os.Exit(0)
}

func maxArea([]int) (_ int) { return }

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
