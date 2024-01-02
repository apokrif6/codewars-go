//Write a program that, given a word, computes the scrabble score for that word.
//Letter Values

//You'll need these:
//Letter                           Value
//A, E, I, O, U, L, N, R, S, T       1
//D, G                               2
//B, C, M, P                         3
//F, H, V, W, Y                      4
//K                                  5
//J, X                               8
//Q, Z                               10

//There will be a preloaded map DictScores with all these values: DictScores["E"] == 1

//Empty string should return 0. The string can contain spaces and letters (upper and lower case), you should calculate the scrabble score only of the letters in that string.
//""           --> 0
//"STREET"    --> 6
//"st re et"    --> 6
//"ca bba g  e" --> 14

package kata

import "strings"

func ScrabbleScore(st string) int {
  stcap:= strings.ToUpper(st);

  var score int;
  
  for _, i := range stcap {
    score += DictScores[string(i)];
  }

  return score
}
