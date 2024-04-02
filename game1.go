package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Increment the wait group counter for each goroutine
	wg.Add(3)

	// 0 is not your turn but the game is up. 1 is waiting for your turn,
	// 2 the game ends you won the game, and 3 the game ends and you lost.
	comunicationP := 0 // controler uses this to comunicate with police.
	comunicationT := 0 // controler uses this to comunicate with thieft.

	var coordinatesP [2]int // police coordinates
	var coordinatesT [2]int // thiets coordinates

	var x int // x size of grid
	var y int // y size of grid
	var s int // s amount of rounds

	x = rand.Intn(200-10) + 10
	y = rand.Intn(200-10) + 10
	s = rand.Intn(10*int(math.Max(float64(x), float64(y)))-2*int(math.Max(float64(x), float64(y)))) + 2*int(math.Max(float64(x), float64(y)))

	coordinatesP[0] = 0
	coordinatesP[1] = 0
	coordinatesT[0] = x - 1
	coordinatesT[1] = y - 1

	gameUp := true  // game runs if game is up
	played := false // for players let the controler know when they played

	fmt.Println("Game")

	go func() {
		// controler
		fmt.Println("Game Started")
		rounds := 0 // amounts of rounds passed

		comunicationP = 1

		for gameUp {
			if coordinatesP[0] == 0 && coordinatesP[1] == 0 && coordinatesT[0] == 0 && coordinatesT[1] == 0 {
				// check end game cases
				gameUp = false
				comunicationP = 3
				comunicationT = 3
				fmt.Println("The game ends at a tie")
			} else if coordinatesP[0] == coordinatesT[0] && coordinatesP[1] == coordinatesT[1] {
				gameUp = false
				comunicationP = 2
				comunicationT = 3
				fmt.Println("The police cought the thieft at: ")
			} else if coordinatesT[1] == 0 || rounds == s {
				gameUp = false
				comunicationP = 3
				comunicationT = 2
				fmt.Println("Theift won.")
			} else if coordinatesP[0] < 0 {
				// Wall setup
				coordinatesP[0] += 1
				played = !played
			} else if coordinatesP[0] == x {
				coordinatesP[0] -= 1
				played = !played
			} else if coordinatesT[0] < 0 {
				coordinatesT[0] += 1
				played = !played
			} else if coordinatesT[0] == x {
				coordinatesT[0] -= 1
				played = !played
			} else if coordinatesP[1] < 0 {
				coordinatesP[1] += 1
				played = !played
			} else if coordinatesP[1] == y {
				coordinatesP[1] -= 1
				played = !played
			} else if coordinatesT[1] < 0 {
				coordinatesT[1] += 1
				played = !played
			} else if coordinatesT[1] == y {
				coordinatesT[1] -= 1
				played = !played
			} else if played {
				// next round
				rounds += 1
				if comunicationP == 0 {
					comunicationP = 1
					comunicationT = 0
				} else {
					comunicationP = 0
					comunicationT = 1
				}
				played = !played
			}
		}
		defer wg.Done()
	}()

	go func() {
		// police

		fmt.Println("Police started")

		for gameUp {
			if comunicationP == 1 && !played {
				direction := rand.Intn(4)

				if direction == 0 {
					coordinatesP[0] += 1
				} else if direction == 1 {
					coordinatesP[0] -= 1
				} else if direction == 2 {
					coordinatesP[1] += 1
				} else if direction == 3 {
					coordinatesP[1] -= 1
				}
				fmt.Println(coordinatesP[0], coordinatesP[1])

				played = !played
			}
		}
		defer wg.Done()
	}()

	go func() {
		// thieft

		fmt.Println("Theift started")

		for gameUp {
			if comunicationT == 1 && !played {
				direction := rand.Intn(4)

				if direction == 0 {
					coordinatesT[0] += 1
				} else if direction == 1 {
					coordinatesT[0] -= 1
				} else if direction == 2 {
					coordinatesT[1] += 1
				} else if direction == 3 {
					coordinatesT[1] -= 1
				}
				fmt.Println(coordinatesT[0], coordinatesT[1])

				played = !played
			}
		}
		defer wg.Done()
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
