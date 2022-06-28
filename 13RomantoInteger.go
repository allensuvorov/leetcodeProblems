// 1354. Construct Target Array With Multiple Sums
// ideas: recursion, will it help with time complexity?
// ideas: heap
package main

import (
	"fmt"
  "reflect"
)

func romanToInt(s string) int {
  
  m := map[rune]int{
    73: 1,    // I
    86: 5,    // V
    88: 10,   // X
    76: 50,   // L
    67: 100,  // C
    68: 500,  // D
    77: 1000, // M
  }

  // sRune := []rune(s)
  
  result := 0
  for _, char := range s {
    fmt.Println(reflect.TypeOf(char))
  }
  
  // for i, r := range s {
  //   fmt.Println(i,r,m[r])
    
  //   if !(m[r] < m[rune(s[(i+1)])]) {
  //     result += m[r]
  //   } else{
  //     result -= m[r]
  //   }
  // }

  for i := 0; i < len(s); i++ {  
    if (!(m[s[i]]<m[s[i+1]])) {    
      result += m[s[i]]; 
    } else {
      result -= m[s[i]];    
    } 
  }


  return result
}

func main() {
  romanToInt("IVXLCDM") //'I', 'V', 'X', 'L', 'C', 'D', 'M'
}
