package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
     
  var (
    sum [] byte
    byte1, byte2, carry byte
    i1 int = len(num1)-1
    i2 int = len(num2)-1
  )
  
  for i1 >=0 || i2 >=0 {
    byte1, byte2 = 0, 0
    
    if i1 >= 0 {
      byte1 = ((num1[i1]-'0'))
      i1 --
    }
    if i2 >= 0 {
      byte2 = ((num2[i2]-'0'))
      i2 -- 
    }
    
    if (byte1+byte2+carry) < 10 { // shifting to bytes
      sum = append([]byte{byte1+byte2+carry+'0'}, sum...)
      carry = 0
    } else {
      sum = append([]byte{(byte1+byte2+carry)%10+'0'}, sum...)
      carry = 1
    }
    // fmt.Print(sum)
  }
  
  if carry == 1 {
    sum = append([]byte{1+'0'}, sum...)
  }
  fmt.Println(string(sum))
  return string(sum)
}

func main() {
  fmt.Println(addStrings("584", "18"))
  fmt.Println(addStrings("9", "1"))
}

// string -> bytes -> string
  // 012345678 +
  //      9999
  //           =
  //         7
  //      
