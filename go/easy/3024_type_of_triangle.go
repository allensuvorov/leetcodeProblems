func triangleType(nums []int) string {
    a, b, c := nums[0], nums[1], nums[2]
    if !(a < b + c && b < a + c && c < b + a) {
        return "none"
    }
    if a == b && b == c {
        return "equilateral"
    }
    if a == b || b == c || c == a {
        return "isosceles"
    }
    return "scalene"
}
