package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS, mapT := makeMap(s), makeMap(t)
	for k, v := range mapS {
		if v != mapT[k] {
			return false
		}
	}
	return true
}

func makeMap(s string) map[rune]int {
	var m map[rune]int = make(map[rune]int)

	for _, v := range s {
		m[v]++
	}
	return m
}
