func numRescueBoats(people []int, limit int) int {

    count := 0
    sort.Ints(people)
    
    for i := len(people) - 1; i >=0; i -- {
        if people[i] == 0 {
            continue
        }

        dif := limit - people [i]
        count++
        if dif == 0 {
            continue
        }
        
        // search via iteration
        for j := i - 1; j >= 0; j-- {
            if people[j] == 0 {
                continue
            }

            if people[j] <= dif{
                people[j] = 0
                break
            }
        }
    }
    return count
}
