func lemonadeChange(bills []int) bool {
    m := map[int]int{}
    for _, v := range bills {
        m[v]++
        switch v {
        case 10:
            if m[5] == 0 {
                return false
            } else {
                m[5]--
            }
        case 20: 
            if m[5] < 3 && (m[5]==0 || m[10]==0) {
                return false
            } else {
                if m[10] > 0{
                    m[10]--
                    m[5]--
                } else {
                    m[5] -= 3
                }
            }
        }
    }
    return true
}
