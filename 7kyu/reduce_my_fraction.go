//Write a function which reduces fractions to their simplest form! Fractions will be presented as an array/tuple (depending on the language) of strictly positive integers, and the reduced fraction must be returned as an array/tuple:

//input:   [numerator, denominator]
//output:  [reduced numerator, reduced denominator]
//example: [45, 120] --> [3, 8]

//All numerators and denominators will be positive integers.

package kata

func ReduceFraction(fraction [2]int) [2]int {
  g := gcd(fraction[0], fraction[1])

  return [2]int{fraction[0]/g, fraction[1]/g}
}
func gcd(a, b int) int{
  for b != 0 {
    a, b = b, a % b
  }
  
  return a
}
