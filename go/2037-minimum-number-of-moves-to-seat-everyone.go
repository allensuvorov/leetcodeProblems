func minMovesToSeat(seats []int, students []int) int {
    seatsCount := inventory(seats)
    studentsCount := inventory(students)
    var i, j int
    result := 0
    for i < len(seatsCount) && j < len(seatsCount)  {
        if seatsCount[i] == 0 {
            i++
            continue
        }
        if studentsCount[j] == 0 {
            j++
            continue
        }
        result += absDif(i, j)
        seatsCount[i] --
        studentsCount[j] --
    }
    return result
}

func absDif(i,j int) int{
    if i > j {
        return i - j
    } else {
        return j - i
    }
}

func inventory(values []int) [101]int {
        result := [101]int{}
        for i := range values {
            result[values[i]]++
        }
    return result
}
