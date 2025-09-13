func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    landWater := earliestFinishTimeInOrder(landStartTime, landDuration, waterStartTime, waterDuration)
    waterLand := earliestFinishTimeInOrder(waterStartTime, waterDuration, landStartTime, landDuration)
    return min(landWater, waterLand)
}

func earliestFinishTimeInOrder(start1, duration1, start2, duration2 []int) int {
    earliestFinishTime1 := math.MaxInt
    for i := range start1 {
        earliestFinishTime1 = min(earliestFinishTime1, start1[i] + duration1[i])
    }
    result := math.MaxInt
    for i := range start2 {
        result = min(result, max(earliestFinishTime1, start2[i])+ duration2[i])
    }
    return result
}
