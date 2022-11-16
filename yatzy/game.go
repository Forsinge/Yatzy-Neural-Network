package yatzy

import (
	"fmt"
	"math/rand"

	input "github.com/Forsinge/Yatzy-Neural-Network/input"
)

func RunGame() {
	fmt.Println("Welcome to Yatzy!")

	sheet := NewSheet()

	for sheet.Categories_left > 0 {
		fmt.Println()
		fmt.Println("----- NEW ROUND -----")
		fmt.Println()
		fmt.Println("Current sheet:")
		for i := range sheet.Categories {
			category := sheet.Categories[i]
			if category.Used && category.Points == 0 {
				fmt.Print("[-] ")
			} else if category.Points == 0 {
				fmt.Print("[ ] ")
			} else {
				fmt.Print("[", category.Points, "] ")
			}

			fmt.Println(category.Name)
		}
		fmt.Println()

		dice := Roll()
		fmt.Println("You rolled:", dice.Faces)
		fmt.Println("Now enter your rerolls. You have 2 rolls left.")

		dice = Reroll(dice, input.ScanInts())
		fmt.Println("You rolled:", dice.Faces)
		fmt.Println("Now enter your rerolls. You have 1 rolls left.")

		dice = Reroll(dice, input.ScanInts())
		fmt.Println("You rolled:", dice.Faces)

		options := ValidCategories(sheet, dice)
		bust := len(options) == 0
		if bust {
			fmt.Println("No options. Select category to discard.")
			options = ActiveCategories(sheet)
		} else {
			fmt.Println("Select a category.")
		}

		for i := range options {
			category := sheet.Categories[options[i]]
			fmt.Println(category.index, category.Name)
		}

		index := input.ScanInt()
		sheet = ConsumeCategory(sheet, dice, index)
	}

	fmt.Println("Game complete! Total score:", GameScore(sheet))
}

func RunRandomizedGame() int {
	sheet := NewSheet()

	for sheet.Categories_left > 0 {
		dice := Roll()
		dice = Reroll(dice, input.RandomRerolls())
		dice = Reroll(dice, input.RandomRerolls())

		options := ValidCategories(sheet, dice)
		bust := len(options) == 0

		if bust {
			options = ActiveCategories(sheet)
		}

		index := options[rand.Intn(len(options))] // because their indices are user adjusted

		sheet = ConsumeCategory(sheet, dice, index)
	}

	/*
		fmt.Println("Final sheet:")
		for i := range sheet.categories {
			category := sheet.categories[i]
			if category.used && category.points == 0 {
				fmt.Print("[-] ")
			} else if category.points == 0 {
				fmt.Print("[ ] ")
			} else {
				fmt.Print("[", category.points, "] ")
			}

			fmt.Println(category.Name)
		}
		fmt.Println()
		fmt.Println("Total score:", GameScore(sheet))
	*/
	return GameScore(sheet)
}
