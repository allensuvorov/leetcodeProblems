type NumArray struct {
    prefSums []int
}


func Constructor(nums []int) NumArray {
    prefSums := make([]int, len(nums))
    prefSum := 0
    for i, v := range nums {
        prefSum += v
        prefSums[i] = prefSum
    }
    return NumArray{prefSums}
}


func (this *NumArray) SumRange(left int, right int) int {
    rangeSum := this.prefSums[right]
    if left > 0 {
        rangeSum = rangeSum - this.prefSums[left - 1]
    }
    return rangeSum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
