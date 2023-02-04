func uniqueMorseRepresentations(words []string) int {
    morseArray := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    englishMorseMap := map[byte]string{}
    for i, v := range morseArray {
        englishMorseMap[byte(i)+'a']= v
    }
    transfMap := map[string]int{}
    for _, word := range words {
        t := ""
        for i := range word {
            t = t + englishMorseMap[word[i]]
        }
        transfMap[t]++
    }
    return len(transfMap)
}
