// working on a bitwise solution
func singleNumber(nums []int) int {
  var ones int = 0
	var twos int = 0

	for i := 0; i < len(nums); i++ {
		var number int = nums[i]
		ones ^= (number & ^twos)
		twos ^= (number & ^ones)
	}

	return ones
}
