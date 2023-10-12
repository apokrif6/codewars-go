//Create a function add(n)/Add(n) which returns a function that always adds n to any number

//var addOne = Add(1)
//addOne(3) // 4

package kata

func Add(n int) func(int)int {
  return func(x int) int {
    return x + n
  }
}
