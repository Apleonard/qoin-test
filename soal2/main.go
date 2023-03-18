package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id    int
	score int
	dice  []int
	next  *Player
}

func main() {
	var numPlayers, numDice = 4, 4

	players := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &Player{id: i + 1, score: 0, dice: make([]int, numDice)}
	}
	for i := 0; i < numPlayers; i++ {
		players[i].next = players[(i+1)%numPlayers]
	}

	round := 1
	for {
		fmt.Printf("Round %d lempar dadu:\n", round)
		for _, player := range players {
			if len(player.dice) > 0 {
				player.rollDice()
				fmt.Printf("Pemain #%d (%d): %v\n", player.id, player.score, player.dice)
			}
		}
		fmt.Println("Setelah evaluasi:")
		activePlayers := 0
		for _, player := range players {
			if len(player.dice) > 0 {
				activePlayers++
				player.evaluateDice()
				fmt.Printf("Pemain #%d (%d): %v\n", player.id, player.score, player.dice)
			}
		}
		if activePlayers <= 1 {
			break
		}
		round++
	}

	var winner *Player
	for _, player := range players {
		if winner == nil || player.score > winner.score {
			winner = player
		}
	}

	fmt.Println("Game berakhir karena hanya pemain #", winner.id, "yang memiliki dadu.")
	fmt.Println("Game dimenangkan oleh pemain #", winner.id, "dengan score", winner.score, "!")
}

func (p *Player) rollDice() {
	rand.Seed(time.Now().UnixNano())
	for i := range p.dice {
		p.dice[i] = rand.Intn(6) + 1
	}
}

func (p *Player) evaluateDice() {
	numDice := len(p.dice)
	for i := 0; i < numDice; i++ {
		if p.dice[i] == 6 {
			p.score++
			p.dice = append(p.dice[:i], p.dice[i+1:]...)
			numDice--
			i--
		} else if p.dice[i] == 1 {
			p.next.dice = append(p.next.dice, 1)
			p.dice = append(p.dice[:i], p.dice[i+1:]...)
			numDice--
			i--
		}
	}
}
