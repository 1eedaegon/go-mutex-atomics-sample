package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
}

func (p *Player) GetHealth() int {
	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) TakeDamage(value int) int {
	atomic.StoreInt32(&p.health, int32(p.GetHealth()-value))
	return int(atomic.LoadInt32(&p.health))
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

// Print player health
func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("Player health: %d\n", p.GetHealth())
		<-ticker.C
	}
}

// Interact player's health
func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.TakeDamage(rand.Intn(40))
		if p.GetHealth() <= 0 {
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
