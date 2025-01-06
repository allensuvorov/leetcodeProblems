// one que
func predictPartyVictory(senate string) string {
    q := []byte(senate)
    r, d := 0, 0
    for _, v := range senate {
        if v == 'R' {
            r++
        } else {
            d++
        }
    }

    bal := 0
    for r != 0 && d != 0 {
        if bal > 0 {
            if q[0] == 'R' {
                bal++
                q = append(q, q[0])
            } else {
                bal--
                d--
            }
            q = q[1:]
        } else if bal < 0 {
            if q[0] == 'D' {
                bal--
                q = append(q, q[0])
            } else {
                bal++
                r--
            }
            q = q[1:]
        } else {
            if q[0] == 'R' {
               bal++
            } else {
                bal--
            }
            q = append(q, q[0])
            q = q[1:]
        }        
    }

    if r > 0 {
        return "Radiant"
    }
    return "Dire"
}

// 2 queues implemented via channels
func predictPartyVictory(senate string) string {
    n := len(senate)
    r, d := make(chan int, n), make(chan int, n)

    for i, c := range senate {
        if c == 'R' {
            r <- i
        } else {
            d <- i
        }
    }

    for len(r) > 0 && len(d) > 0 {
        rHead := <- r
        dHead := <- d
        if rHead < dHead {
            r <- rHead + n
        } else {
            d <- dHead + n
        }
    }

    if len(r) == 0 { return "Dire" }
    return "Radiant"
}

// 2 queues implemented via slices
func predictPartyVictory(senate string) string {
    n := len(senate)
    r, d := []int{}, []int{}
    for i, c := range senate {
        if c == 'R' {
            r = append(r, i)
        } else {
            d = append(d, i)
        }
    }
    for len(r) > 0 && len(d) > 0 {
        rHead := r[0]
        r = r[1:]
        dHead := d[0]
        d = d[1:]
        if rHead < dHead {
            r = append(r, rHead + n)
        } else {
            d = append(d, dHead + n)
        }
    }
    if len(r) == 0 { return "Dire" }
    return "Radiant"
}
