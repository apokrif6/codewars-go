//Given a string, turn each character into its ASCII character code and join them together to create a number - let's call this number total1:

//'ABC' --> 'A' = 65, 'B' = 66, 'C' = 67 --> 656667

//Then replace any incidence of the number 7 with the number 1, and call this number 'total2':

//total1 = 656667
              ^
//total2 = 656661
              ^

//Then return the difference between the sum of the digits in total1 and total2:

//  (6 + 5 + 6 + 6 + 6 + 7)
//- (6 + 5 + 6 + 6 + 6 + 1)
//-------------------------
//                       6

package kata

import (
  "strconv";
  "strings";
)

func Calc(s string) int {
  var sum string
  var sumOfDigits int
  var sumOfDigitsReplaced int
  
  for _, element := range s {
    sum += strconv.Itoa(int(element))
  }
  
  replaced := strings.ReplaceAll(sum, "7", "1")
  
  for _, char := range sum {
    sumOfDigits += int(char - '0')
  }
  
  for _, char := range replaced {
    sumOfDigitsReplaced += int(char - '0')
  }
  
  return sumOfDigits - sumOfDigitsReplaced
}
