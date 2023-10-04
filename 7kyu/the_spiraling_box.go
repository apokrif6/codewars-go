//Given two positive integers m (width) and n (height), fill a two-dimensional list (or array) of size m-by-n in the following way:

//    (1) All the elements in the first and last row and column are 1.
//    (2) All the elements in the second and second-last row and column are 2, excluding the elements from step 1.
//    (3) All the elements in the third and third-last row and column are 3, excluding the elements from the previous steps.

//    And so on ...

//Examples

//Given m = 5, n = 8, your function should return

//[[1, 1, 1, 1, 1],
// [1, 2, 2, 2, 1],
// [1, 2, 3, 2, 1],
// [1, 2, 3, 2, 1],
// [1, 2, 3, 2, 1], 
// [1, 2, 3, 2, 1],
// [1, 2, 2, 2, 1],
// [1, 1, 1, 1, 1]]

//Given m = 10, n = 9, your function should return

//[[1, 1, 1, 1, 1, 1, 1, 1, 1, 1],
// [1, 2, 2, 2, 2, 2, 2, 2, 2, 1],
// [1, 2, 3, 3, 3, 3, 3, 3, 2, 1], 
// [1, 2, 3, 4, 4, 4, 4, 3, 2, 1], 
// [1, 2, 3, 4, 5, 5, 4, 3, 2, 1], 
// [1, 2, 3, 4, 4, 4, 4, 3, 2, 1], 
// [1, 2, 3, 3, 3, 3, 3, 3, 2, 1], 
// [1, 2, 2, 2, 2, 2, 2, 2, 2, 1], 
// [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]]

package kata

func CreateBox(width, length int) [][]int {
  matrix := make([][]int, length)
  for i := range matrix {
      matrix[i] = make([]int, width)
  }

  for layer := 0; layer < (min(width, length) + 1) / 2; layer++ {
      num := layer + 1
      for i := layer; i < width - layer; i++ {
          matrix[layer][i] = num
          matrix[length - 1 - layer][i] = num
      }
      for i := layer; i < length - layer; i++ {
          matrix[i][layer] = num
          matrix[i][width - 1 - layer] = num
      }
  }

  return matrix
}

func min(a, b int) int {
  if a < b {
      return a
  }
  return b
}
