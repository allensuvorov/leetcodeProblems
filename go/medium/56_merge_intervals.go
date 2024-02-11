func merge(intervals [][]int) [][]int {
   sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
   merged := [][]int{}

   for _, interval := range intervals {
       last := len(merged) - 1
       if len(merged) == 0 || merged[last][1] < interval[0] {
           merged = append(merged, interval)
       } else {
           merged[last][1] = max(merged[last][1], interval[1])
       }
   }
   return merged
}
