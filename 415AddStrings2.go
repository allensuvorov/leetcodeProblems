// changing longer slice

package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
  if len(num1)<len(num2) {
    num1, num2 = num2, num1
  }

  var (
    res = []byte(num1)
    rem int = len(num1)-len(num2)
    carry byte = 0
    digit byte = 0
    
  )

  for i:=len(res)-1; i > rem-1; i-- {
    
    digit = num1[i] + num2[i-rem] - '0'*2 + carry
    
    if digit < 10 {
      res[i] = digit + '0'
      carry = 0
    } else {
      res[i] = digit - 10 + '0'
      carry = 1
    }
    fmt.Println("digit is ", digit, "carry is ", carry, "res is ", string(res))
  }

  for i:=rem-1; i >=0; i--{
    if carry == 0 {
      break
    }
    if carry == 1 {
      digit = res[i] - '0'+ carry
      fmt.Println("digit2 is ", digit)
      if digit < 10 {
        res[i] = digit +'0'
        carry = 0
        break
      } else {
        res[i] = '0'
        carry = 1
      }  
    }
  }
  if carry == 1 {
    res = append([]byte{'1'}, res...)
  }
  return string(res)
}

func main() {
  // fmt.Println(addStrings("584", "18"))
  // fmt.Println(addStrings("9", "1"))
  fmt.Println(addStrings("9", "99"))
}
