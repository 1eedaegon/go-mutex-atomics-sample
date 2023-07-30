package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	mu     sync.RWMutex
	health int
}

func (p *Player) GetHealth() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) TackDamage(value int) int {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.health -= value
	return p.health
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
		p.TackDamage(rand.Intn(40))
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
