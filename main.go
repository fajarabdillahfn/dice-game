package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type PlayerDetail struct {
	ID            int
	Point         int
	DiceRemaining int
	ResDice       [][]int
	EvaluatedRes  [][]int
}

func main() {
	var player, dice int

	fmt.Print("masukkan jumlah pemain:")
	_, err := fmt.Scanln(&player)
	if err != nil {
		panic(err)
	}

	fmt.Print("masukkan jumlah dadu:")
	_, err = fmt.Scanln(&dice)
	if err != nil {
		panic(err)
	}

	winners := gameStart(player, dice)

	if len(winners) == 1 {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.", winners[0])
	} else {
		fmt.Print("Game berakhir dengan hasil imbang antara pemain ")
		for i, winner := range winners {
			fmt.Printf("#%d", winner)

			if i == len(winners)-2 {
				fmt.Print(" dan ")
			} else if i < len(winners)-2 {
				fmt.Print(", ")
			}
		}
		fmt.Println(".")
	}
}

func gameStart(n, m int) []int {
	fmt.Printf("Pemain = %d, Dadu = %d\n", n, m)
	fmt.Println("==================")

	var players []*PlayerDetail
	var max int

	// initialize players
	for i := 1; i <= n; i++ {
		players = append(players, &PlayerDetail{
			ID:            i,
			Point:         0,
			DiceRemaining: m,
		})
	}

	// GAME PLAY!!
	round := 1
	for {
		maxPoint, isDone := gamePlay(round, players)

		if isDone {
			max = maxPoint
			break
		}
		round++
	}

	// winner checker
	winners := []int{}
	for _, player := range players {
		if player.Point == max {
			winners = append(winners, player.ID)
		}
	}

	return winners
}

func gamePlay(round int, players []*PlayerDetail) (max int, isDone bool) {
	fmt.Printf("Giliran %d lempar dadu\n", round)
	rand.Seed(time.Now().UnixNano())

	addDice := map[int]int{}
	for _, player := range players {
		if player.DiceRemaining > 0 {
			player.ResDice = append(player.ResDice, []int{})
			player.EvaluatedRes = append(player.EvaluatedRes, []int{})
			remainDice := player.DiceRemaining
			for j := 0; j < remainDice; j++ {
				num := rand.Intn(6) + 1
				player.ResDice[round-1] = append(player.ResDice[round-1], num)

				switch num {
				case 6:
					player.Point++
					player.DiceRemaining--
				case 1:
					player.DiceRemaining--
					if player.ID == len(players) {
						addDice[1]++
					} else {
						addDice[player.ID+1]++
					}
				default:
					player.EvaluatedRes[round-1] = append(player.EvaluatedRes[round-1], num)
				}
			}

			fmt.Printf("\tPemain #%d (%d): %s\n", player.ID, player.Point, strings.Trim(strings.Replace(fmt.Sprint(player.ResDice[round-1]), " ", ",", -1), "[]"))
		} else {
			fmt.Printf("\tPemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.ID, player.Point)
		}
	}

	fmt.Println("Setelah evaluasi:")
	playersRemain := []int{}
	for _, player := range players {
		for i := 0; i < addDice[player.ID]; i++ {
			player.DiceRemaining++
			player.EvaluatedRes[round-1] = append(player.EvaluatedRes[round-1], 1)
		}

		if player.Point > max {
			max = player.Point
		}

		if player.DiceRemaining > 0 {
			fmt.Printf("\tPemain #%d (%d): %s\n", player.ID, player.Point, strings.Trim(strings.Replace(fmt.Sprint(player.EvaluatedRes[round-1]), " ", ",", -1), "[]"))

			playersRemain = append(playersRemain, player.ID)
		} else {
			fmt.Printf("\tPemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.ID, player.Point)
		}
	}
	fmt.Println("==================")
	if len(playersRemain) == 1 {
		fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", playersRemain[0])
		isDone = true
	}

	return
}
