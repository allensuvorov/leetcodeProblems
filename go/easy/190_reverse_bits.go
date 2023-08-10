func reverseBits(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res |= n & 1
		n >>= 1
	}
	return res
}
