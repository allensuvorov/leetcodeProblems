func totalWaviness(num1 int, num2 int) int { // linear
    totalWavecount := 0
    for num := num1; num <= num2; num++ {
        totalWavecount += waveCount(num)
    }
    return totalWavecount
}

func waveCount(num int) int { // log time
    count := 0
    q := []int{}
    for ; num > 0; num /= 10 {
        d := num % 10
        q = append(q, d)
        
        if len(q) == 3 {
            if q[0] < q[1] && q[1] > q[2] {
                count++
            }
            if q[0] > q[1] && q[1] < q[2] {
                count++
            }
            q = q[1:]
        }
        
    }
    return count
}
