type Logger struct {
    lastTimePrinted map[string]int
}


func Constructor() Logger {
    return Logger{make(map[string]int)}
}


func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    var ans bool
    // new or 10 sec passed
    if v, ok := this.lastTimePrinted[message]; !ok || timestamp >= v + 10 {
        ans = true
        this.lastTimePrinted[message] = timestamp
    }
    return ans
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
