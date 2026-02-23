func deckRevealedIncreasing(deck []int) []int {
	// sort in in order
	slices.Sort(deck)
	q := []int{}
	for i := len(deck)-1; i>=0; i--{
		// rotate
		if len(q) > 0 {
			front := q[0]
			q = append(q, front)
			q = q[1:]
		}
		// append
		q = append(q, deck[i])
	}
    slices.Reverse(q)
	return q
}
