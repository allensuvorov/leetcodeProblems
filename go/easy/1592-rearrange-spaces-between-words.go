import "strings"
func reorderSpaces(text string) string {
    words := strings.Fields(text)
    totalSpaces := strings.Count(text, " ")
    carry := 0
    evenSpacesCount := 0

    if len(words) == 1 {
        carry = totalSpaces
    } else {
        carry = totalSpaces % (len(words) - 1)
        evenSpacesCount = totalSpaces / (len(words) - 1)
    }
    evenSpaces := strings.Repeat(" ", evenSpacesCount)
    withSpaces := strings.Join(words, evenSpaces)
    trailingSpaces := strings.Repeat(" ", carry)
    return withSpaces + trailingSpaces
}
