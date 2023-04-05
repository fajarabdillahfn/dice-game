package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type PlayerDetail struct {
	ID            int
	Point         int
	DiceRemaining int
	ResDice       []int
	EvaluatedRes  []int
}

func main() {
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.", gameInit(2, 3))
}

func gameInit(n, m int) int {
	fmt.Printf("Pemain = %d, Dadu = %d\n", n, m)
	fmt.Println("==================")

	var players []*PlayerDetail
	var max, win int

	for i := 1; i <= n; i++ {
		players = append(players, &PlayerDetail{
			ID:            i,
			Point:         0,
			DiceRemaining: m,
		})
	}

	round := 1
	for {
		fmt.Printf("Giliran %d lempar dadu\n", round)

		for _, player := range players {
			player.ResDice = []int{}
			for j := 0; j < player.DiceRemaining; j++ {
				num := rand.Intn(6) + 1
				player.ResDice = append(player.ResDice, num)
			}

			if player.DiceRemaining > 0 {
				fmt.Printf("\tPemain #%d (%d): %s\n", player.ID, player.Point, strings.Trim(strings.Replace(fmt.Sprint(player.ResDice), " ", ",", -1), "[]"))
			} else {
				fmt.Printf("\tPemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.ID, player.Point)
			}
		}

		fmt.Println("Setelah evaluasi:")
		addDice := map[int]int{}
		for _, player := range players {
			player.EvaluatedRes = []int{}
			for _, dice := range player.ResDice {
				switch dice {
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
					player.EvaluatedRes = append(player.EvaluatedRes, dice)
				}
			}
		}

		for _, player := range players {
			for i:=0; i<addDice[player.ID]; i++ {
				player.DiceRemaining++
				player.EvaluatedRes = append(player.EvaluatedRes, 1)
			}

			if player.DiceRemaining > 0 {
				fmt.Printf("\tPemain #%d (%d): %s\n", player.ID, player.Point, strings.Trim(strings.Replace(fmt.Sprint(player.EvaluatedRes), " ", ",", -1), "[]"))
			} else {
				fmt.Printf("\tPemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", player.ID, player.Point)
			}
		}
		fmt.Println("==================")

		playersRemain := []int{}
		for _, player := range players {
			if player.DiceRemaining > 0 {
				playersRemain = append(playersRemain, player.ID)
				max = player.Point
			}
		}
		if len(playersRemain) == 1 {
			fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", playersRemain[0])
			win = playersRemain[0]
			break
		}
		round++
	}

	for _, player := range players {
		if player.Point > max {
			max = player.Point
			win = player.ID
		}
	}

	return win
}
