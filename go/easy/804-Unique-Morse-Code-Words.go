func uniqueMorseRepresentations(words []string) int {
    morseArray := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    englishMorseMap := map[byte]string{}
    for i, v := range morseArray {
        englishMorseMap[byte(i)+'a']= v
    }
    transfMap := map[string]int{}
    for _, word := range words {
        t := []string{}
        for i := range word {
            t = append(t, englishMorseMap[word[i]])
        }
        ts := strings.Join(t, "")
        transfMap[ts]++
    }
    return len(transfMap)
}
