package yatzy

import (
	"fmt"
	"math/rand"
)

type Yatzy struct {
	sheet Sheet
	dice  Dice
}

func NewGame() Yatzy {
	return Yatzy{
		sheet: NewSheet(),
		dice:  Roll(),
	}
}

func (game *Yatzy) Print() {
	fmt.Println("___________________________")
	fmt.Println()
	game.dice.Print()
	game.sheet.Print()
}

func NewRollState(game Yatzy) Yatzy {
	if game.sheet.Categories_left <= 0 {
		game.sheet = NewSheet()
		game.dice = Roll()
	}

	if game.dice.Rerolls_left <= 0 {
		cats := validCategories(game.sheet, game.dice)
		if len(cats) == 0 {
			cats = activeCategories(game.sheet)
		}
		i := cats[rand.Intn(len(cats)-1)]
		game.sheet = ConsumeCategory(game.sheet, game.dice, i)
	} else {
		game.dice = Reroll(game.dice, []int{0, 1, 2, 3, 4})
	}

	return game
}

func NewSheetState(game Yatzy) Yatzy {
	if game.sheet.Categories_left <= 0 {
		game.sheet = NewSheet()
		game.dice = Roll()
	}

	cats := validCategories(game.sheet, game.dice)
	if len(cats) == 0 {
		cats = activeCategories(game.sheet)
	}
	i := cats[rand.Intn(len(cats)-1)]
	game.sheet = ConsumeCategory(game.sheet, game.dice, i)

	game.dice = Roll()

	return game
}
