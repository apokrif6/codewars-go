//Task

//Given a string str, reverse it and omit all non-alphabetic characters.
//Example

//For str = "krishan", the output should be "nahsirk".
//For str = "ultr53o?n", the output should be "nortlu".

//Input/Output
//    [input] string str

//A string consists of lowercase latin letters, digits and symbols.
//    [output] a string

package kata

import (
	"regexp"
)

func ReverseLetters(s string) string {
	exp := regexp.MustCompile("[^a-zA-Z]+")
	s = exp.ReplaceAllString(s, "")

	output := []rune(s)
	length := len(output)

	for i := 0; i < length/2; i++ {
		output[i], output[length-i-1] = output[length-i-1], output[i]
	}

	return string(output)
}
