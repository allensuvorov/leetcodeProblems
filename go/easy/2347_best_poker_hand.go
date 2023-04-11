package easy

func bestHand(ranks []int, suits []byte) string {
    isFlush := true
    
    for i := 1; i < len(suits); i++ {
        if suits[0] != suits[i] {
            isFlush = false
        }
    }
    if isFlush {
        return "Flush"
    }

    rm := map[int]int{}

    for i := range ranks {
        rm[ranks[i]]++
    }
    
    maxRankCount := 0
    for _, rc := range rm {
        if rc > maxRankCount {
            maxRankCount = rc
        }
    }

    switch  {
        case maxRankCount >= 3:
            return "Three of a Kind"
        case maxRankCount == 2:
            return "Pair"
        default:
            return "High Card"
    }

}
