//This Kata is intended as a small challenge for my students

//All Star Code Challenge #1

//Write a function, called sumPPG, that takes two NBA player objects/struct/Hash/Dict/Class and sums their PPG

package kata

type NBAPlayer struct {
  Team string
  Ppg  float64
}

func SumPpg(playerOne, playerTwo NBAPlayer) float64 {
  return playerOne.Ppg + playerTwo.Ppg
}
