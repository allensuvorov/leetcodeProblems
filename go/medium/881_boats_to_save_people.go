func numRescueBoats(people []int, limit int) int {
    n := 0
    sort.Ints(people)
    l := 0
    r := len(people) - 1
    for l <= r {
        n++
        if people[l] + people[r] <= limit{
            l++
        }
        r--
    }
    return n
}
