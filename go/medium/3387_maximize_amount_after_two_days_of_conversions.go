func maxAmount(
    initialCurrency string, 
    pairs1 [][]string, rates1 []float64, 
    pairs2 [][]string, rates2 []float64,
    ) float64 {

    pathRates1 := pathRates(initialCurrency, pairs1, rates1)
    pathRates2 := pathRates(initialCurrency, pairs2, rates2)

    res := 1.0

    for currency1, rate1 := range pathRates1 {
        if rate2, ok := pathRates2[currency1]; ok {
            res = max(res, rate1 / rate2)
        }
    }

    return res
}

type node struct {
    currency string
    rate float64
}

func pathRates(initialCurrency string, pairs [][]string, rates []float64) map[string]float64 {
    res := make(map[string]float64)

    q := []node{node{initialCurrency, 1}}
    seen := map[string]bool{initialCurrency:true}
    for len(q) > 0 {
        n := len(q)
        for range n {
            now := q[0]
            q = q[1:]
            // find neighbors
            for i, pair := range pairs {
                if pair[0] == now.currency && !seen[pair[1]] {
                    seen[pair[1]] = true
                    pathRate := now.rate * rates[i]
                    q = append(q, node{pair[1], pathRate})
                    res[pair[1]] = pathRate
                }
                if pair[1] == now.currency && !seen[pair[0]] {
                    seen[pair[0]] = true
                    pathRate := now.rate / rates[i]
                    q = append(q, node{pair[0], pathRate})
                    res[pair[0]] = pathRate
                }
            }
        }
    }

    return res
}
