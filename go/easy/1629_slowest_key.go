func slowestKey(releaseTimes []int, keysPressed string) byte {
    maxDur := 0
    var maxKey byte
    prevTime := 0
    for i, t := range releaseTimes{
        curDur := t - prevTime
        if curDur > maxDur {
            maxDur = curDur
            maxKey = keysPressed[i]
        } else if curDur == maxDur && keysPressed[i] > maxKey {
            maxKey = keysPressed[i]
        }
        prevTime = t
    }
    return maxKey
}
