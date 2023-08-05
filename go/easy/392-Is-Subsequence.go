func isSubsequence(s string, t string) bool {
   track, scout := 0, 0
   for track < len(s) && scout < len(t) {
       if s[track] == t[scout] {
           track++
       }
       scout++
   }
   return track == len(s)
}
