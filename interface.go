// Interfaces explained:
//   -> an interface is a list of functions (behavior)
//   -> any type that has those functions automatically matches the interface
//   -> Go does this implicitly (no "implements" keyword)
//
// Why:
//   -> lets us treat different types the same way
//   -> we don't care WHAT the thing is, only WHAT IT CAN DO
//
// logic:
//   damageable -> interface (anything that can take damage)
//   player / enemy -> structs
//   both have takeDamage()
//   dealDamage() works on BOTH without knowing the concrete type

package main

import "fmt"

// interface: anything that can take damage
type damageable interface {
	takeDamage(amount int)
}

// concrete type #1
type player struct {
	health int
}

// player satisfies damageable
func (p *player) takeDamage(amount int) {
	p.health -= amount
}

// concrete type #2
type enemy struct {
	health int
}

// enemy also satisfies damageable
func (e *enemy) takeDamage(amount int) {
	e.health -= amount
}

// function that does not care what it hits
func dealDamage(d damageable, amount int) {
	d.takeDamage(amount)
}

func main() {
	p := &player{health: 100}
	e := &enemy{health: 50}

	dealDamage(p, 30)
	dealDamage(e, 20)

	fmt.Println("player health:", p.health)
	fmt.Println("enemy health:", e.health)
}
