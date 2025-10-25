func productExceptSelf(nums []int) []int {
    // prefixes ->
    prefixes := make([]int, len(nums))
    prefix := 1
    for i := range nums {
        prefixes[i] = prefix
        prefix = prefix * nums[i]
    }

    // input array - suffixes  <-
    suffix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        prefixes[i] *= suffix
        suffix = suffix * nums[i]
    }

    return prefixes
}
