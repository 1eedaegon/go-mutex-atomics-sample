package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	health int
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

// Print player health
func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("Player health: %d\n", p.health)
		<-ticker.C
	}
}

// Interact player's health
func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.health -= rand.Intn(40)
		if p.health <= 0 {
			fmt.Println("Game Over")
			break
		}
		<-ticker.C
	}
}
func main() {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
