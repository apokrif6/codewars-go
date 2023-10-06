//Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

//Examples:

//input:    output:
//0    ->   0
//2    ->   5
//3    ->   5
//12   ->   15
//21   ->   25
//30   ->   30
//-2   ->   0
//-5   ->   -5
//etc.

package kata

import "math"

func RoundToNext5(n int) int {
  return int(math.Ceil(float64(n)/5.0)) * 5
}
