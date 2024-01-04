//Given an integer n and two other values, build an array of size n filled with these two values alternating.
//Examples

//5, true, false     -->  [true, false, true, false, true]
//10, "blue", "red"  -->  ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
//0, "one", "two"    -->  []

package kata


func Alternate(n int, firstValue string, secondValue string) (alternated []string){
  for i := 0; i < n; i++ {
		if i & 1 == 0 {
			alternated = append(alternated, firstValue);
		} else {
			alternated = append(alternated, secondValue);
		}
	}
  return alternated;
}
