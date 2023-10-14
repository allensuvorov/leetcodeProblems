func findClosestElements(arr []int, k int, x int) []int {
    left := 0
    right := len(arr)-k
    for left < right {
        mid := left + (right-left)/2
        startDist := x-arr[mid]
        endDist := arr[mid+k]-x
        if startDist > endDist {
            left = mid+1
        } else {
            right = mid
        }
    }
    return arr[left:left+k]
}
