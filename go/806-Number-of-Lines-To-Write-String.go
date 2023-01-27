// someone else did a short one 
func numberOfLines(widths []int, s string) []int {
    line:=0
    sum:=0
    for _,val:=range s{
        index:=int(val-'a')
        width:=widths[index]
        if sum+width>100{
            line++
            sum=width
        }else{
            sum+=width
        }
    }
    return []int{line+1,sum}
}
// i did a long one
func numberOfLines(widths []int, s string) []int {
    var sum int
    count := 0
    for i := 0; ; {
        w := widths[rune(s[i]) - 'a']
        
        if sum + w < 100 {
            sum += w
            i++
            if i == len(s) {
                count++
                break
            }
            continue
        }

        if sum + w == 100 {
            sum += w
            count++
            i++
            if i == len(s) {
                break
            }
            sum = 0
            continue
        }

        if sum + w > 100 {
            count++
            sum = 0
        }
    }
    return []int{count, sum}
}
