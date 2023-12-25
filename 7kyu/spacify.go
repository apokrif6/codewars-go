//Modify the spacify function so that it returns the given string with spaces inserted between each character.

//spacify=("hello world") -> # returns "h e l l o   w o r l d

package kata

import "strings"

func Spacify(s string) string {
  return strings.Join(strings.Split(s, ""), " ")
}
