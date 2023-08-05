func isSubsequence(s string, t string) bool {
   collector := 0
   scout := 0
   for collector < len(s) && scout < len(t) {
       if s[collector] == t[scout] {
           collector++
       }
       scout++
   }
   if collector != len(s) {
       return false
   }
   return true
}
