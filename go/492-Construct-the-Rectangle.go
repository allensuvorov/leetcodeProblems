func constructRectangle(area int) []int {
    min := area
    res := make([]int, 2)
    for w := 1; w <= area / w ; w++ {
        if area%w == 0 && area/w - w < min {
            min = area/w - w
            res[0], res[1] = area/w, w
        }
    }
    return res
}
