//Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

//Return your answer as a number.

package kata

import "strconv"

func SumMix(arr []any) int {
  var sum int;
  
  for _, element := range arr {
    switch v := element.(type) {
      case int:
        sum += int(v);
      case string:
        intValue, _ := strconv.Atoi(v);
        sum += intValue;
    }
  }
  
  return sum;
}
