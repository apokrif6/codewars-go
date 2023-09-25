//Write reverseList function that simply reverses lists.

package kata

func ReverseList(lst []int) []int {
  var reversed []int
  
  for i := len(lst) - 1; i >= 0; i-- {
    reversed = append(reversed, lst[i])
  }
  
  return reversed
}
