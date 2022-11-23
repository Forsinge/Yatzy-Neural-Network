package yatzy

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Faces               []int
	Counts              [6]int
	sums                [6]int
	total_sum           int
	group_count         int
	largest_group       int
	largest_group_index int
	Rerolls_left        int
}

func createDice(rolls []int, rerolls_left int) Dice {
	var dice Dice

	for i := range rolls {
		face := rolls[i]
		dice.Counts[face-1] += 1
		dice.sums[face-1] += face
		dice.total_sum += face
	}

	largest_group_index := 0
	largest_group := 0
	group_count := 0
	for i := range dice.Counts {
		count := dice.Counts[i]
		if count != 0 {
			group_count += 1
			if count > largest_group {
				largest_group = count
				largest_group_index = i
			}
		}
	}

	dice.Faces = rolls
	dice.group_count = group_count
	dice.largest_group = largest_group
	dice.largest_group_index = largest_group_index
	dice.Rerolls_left = 2
	return dice
}

func (dice *Dice) Print() {
	fmt.Println("Current dice:")
	fmt.Println(dice.Faces)
	fmt.Println()
}

func Reroll(dice Dice, indices []int) Dice {
	rolls := dice.Faces
	for i := range indices {
		rolls[indices[i]] = rand.Intn(6) + 1
	}

	return createDice(rolls, dice.Rerolls_left-1)
}

func Roll() Dice {
	rolls := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}

	return createDice(rolls, 2)
}
