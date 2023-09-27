//A triangle is called an equable triangle if its area equals its perimeter. Return true, if it is an equable triangle, else return false. You will be provided with the length of sides of the triangle. Happy Coding!

package kata

import "math"

func EquableTriangle(a, b, c int) bool {
  perimeter := float64(a + b + c)
  
  s :=  perimeter/ 2
  area := math.Sqrt(s * (s-float64(a)) * (s-float64(b)) * (s-float64(c)))
  
  return perimeter == area
}
