//A squared string is a string of n lines, each substring being n characters long. We are given two n-squared strings. For example:

//s1 = "abcd\nefgh\nijkl\nmnop" s2 = "qrst\nuvwx\nyz12\n3456"

//Let us build a new string strng of size (n + 1) x n in the following way:

//    The first line of strng has the first char of the first line of s1 plus the chars of the last line of s2.
//    The second line of strng has the first two chars of the second line of s1 plus the chars of the penultimate line of s2 except the last char
//    and so on until the nth line of strng has the n chars of the nth line of s1 plus the first char of the first line of s2.

//Calling this function compose(s1, s2) we have:

//compose(s1, s2) -> "a3456\nefyz1\nijkuv\nmnopq"
//or printed:
//abcd    qrst  -->  a3456
//efgh    uvwx       efyz1
//ijkl    yz12       ijkuv
//mnop    3456       mnopq

package kata

import "strings"

func Compose(s1, s2 string) string {
  lines1 := strings.Split(s1, "\n")
  lines2 := strings.Split(s2, "\n")
  n := len(lines1)

  var result strings.Builder

  for i := 0; i < n; i++ {
      result.WriteString(lines1[i][:i+1])
      result.WriteString(lines2[n-i-1][:n-i])
      if i < n-1 {
          result.WriteString("\n")
      }
  }

  return result.String()
}
