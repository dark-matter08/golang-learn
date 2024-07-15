//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health struct {
	value uint
	max   uint
}

type Energy struct {
	value uint
	max   uint
}

type Player struct {
	health Health
	energy Energy
	name   string
}

func (player *Player) addHealth(increment uint) {
	fmt.Println("Increasing player health by", increment)

	if player.health.value+increment >= player.health.max {
		player.health.value = player.health.max
		fmt.Println("Player stats after change:", player)

		return
	}

	player.health.value += increment
	fmt.Println("Player stats after change:", player)

}

func (player *Player) addEnergy(increment uint) {
	fmt.Println("Increasing player Energy by", increment)

	if player.energy.value+increment >= player.energy.max {
		player.energy.value = player.energy.max
		fmt.Println("Player stats after change:", player)

		return
	}

	player.energy.value += increment
	fmt.Println("Player stats after change:", player)

}

func (player *Player) reduceHealth(decrement uint) {
	fmt.Println("Decreasing player health by", decrement)

	if player.health.value-decrement > player.health.value {
		player.health.value = 0
		fmt.Println("Player stats after change:", player)

		return
	}

	player.health.value -= decrement
	fmt.Println("Player stats after change:", player)

}

func (player *Player) reduceEnergy(decrement uint) {
	fmt.Println("Decreasing player energy by", decrement)

	if player.energy.value-decrement > player.energy.value {
		player.energy.value = 0
		fmt.Println("Player stats after change:", player)

		return
	}

	player.energy.value -= decrement
	fmt.Println("Player stats after change:", player)

}

func main() {

	player := Player{name: "Nde Lucien", health: Health{value: 10, max: 15}, energy: Energy{value: 12, max: 15}}

	fmt.Println("Initial Player:", player)

	player.addHealth(1)
	player.addHealth(1)
	player.reduceHealth(5)
	player.addEnergy(3)
	player.reduceEnergy(8)
	player.reduceEnergy(20)

}
