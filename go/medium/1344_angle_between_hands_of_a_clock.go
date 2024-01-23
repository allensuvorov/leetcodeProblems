func angleClock(hour int, minutes int) float64 {
    h, m := float64(hour), float64(minutes)
    var hDeg float64 = h * 30 + m / 2
    var mDeg float64 = m * 6
    absDiff := math.Abs(hDeg - mDeg)
    if absDiff > 180 {
        return 360 - absDiff
    }
    return absDiff
}
