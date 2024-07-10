func minOperations(logs []string) int {
    balance := 0
    for _, v := range logs {
        if v == "../" {
            if balance > 0 {
                balance--
            }
        } else if v != "./" {
            balance++
        }
    }
    return balance
}
