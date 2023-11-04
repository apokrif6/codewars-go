//Create function fib that returns n'th element of Fibonacci sequence (classic programming task).

package kata

func Fib(n int) int {
  if n <= 1 {
    return n;
  }
  
  a, b := 0, 1
  
  for i := 2; i <= n; i++ {
    c := a + b
    a = b
    b = c
  }
  
  return b
}
