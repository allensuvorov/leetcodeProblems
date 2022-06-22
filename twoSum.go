func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		var b int = target - num
		fmt.Println(b)

		if v, found := numMap[b]; found {
			return []int{v, i}
		}
		numMap[num] = i
		fmt.Println(numMap)
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
