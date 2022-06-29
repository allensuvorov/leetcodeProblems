package main

import (
	"fmt"
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

  sRune := []rune(s)

  var (
    result, r, next int
  )
  
  for i := 0; i < len(sRune)-1; i++ {  
    r = m[rune(s[i])]
    next = m[rune(s[i+1])]
    if r < next {    
      result -= r 
    } else {
      result += r 
    }
  }
  
  last := sRune[len(sRune)-1]
  result += m[last]
  return result
}

func main() {
  fmt.Println(romanToInt("IVXLCDM"))
  fmt.Println(romanToInt("IV"))
}
