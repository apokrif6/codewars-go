//Write a function func PickGrains(<-chan string) (int, int), which takes a channel of strings and returns two integers: the first one with the amount of good occurrences, the second one with the amount of bad occurrences.
//Example
//
// channel contains 3 times "good" and 4 times "bad"
//grains := make(chan string, 7)
//grains<-"good"
//grains<-"bad"
//grains<-"bad"
//grains<-"good"
//grains<-"bad"
//grains<-"bad"
//grains<-"good"
//close(grains)

// your implementation of the PickGranes function
//good, bad := PickGrains(grains)
// good must be 3
// bad must be 4

package kata

func PickGrains(grains <-chan string) (good int, bad int) {
  for grain := range grains {
    if grain == "good" {
      good++
    } else if grain == "bad" {
      bad++
    }
  }

  return good, bad
}
