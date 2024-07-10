func minOperations(logs []string) int {
    balance := 0
    for _, v := range logs {
        switch {
        case v == "../": 
            if balance > 0 {
                balance--
            }
        case v != "./":
            balance++
        }
    }
    return balance
}
