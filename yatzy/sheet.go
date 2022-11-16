package yatzy

import (
	"math/rand"
)

type Sheet struct {
	Categories        [15]Category
	Categories_left   int
	upper_section_sum int
	total_sum         int
}

type Category struct {
	Name   string
	Points int
	index  int
	Used   bool
}

func NewSheet() Sheet {
	var sheet Sheet
	category_Names := [15]string{
		"Ones", "Twos", "Threes", "Fours", "Fives", "Sixes", "Pairs",
		"Two Pairs", "Three of a kind", "Four of a kind", "Small straight",
		"Large straight", "Full house", "Chance", "Yatzy"}

	for i := range category_Names {
		category := Category{
			Name:   category_Names[i],
			index:  i, // adjusted for user
			Points: 0,
			Used:   false,
		}
		sheet.Categories[i] = category
	}

	sheet.Categories_left = 15
	return sheet
}

func isValid(category Category, dice Dice, index int) bool {
	isValid := !category.Used
	switch index {
	case 0, 1, 2, 3, 4, 5:
		isValid = isValid && dice.Counts[index] != 0
	case 6:
		isValid = isValid && dice.largest_group >= 2
	case 7:
		isValid = isValid && dice.group_count+dice.largest_group == 5
	case 8:
		isValid = isValid && dice.largest_group >= 3
	case 9:
		isValid = isValid && dice.largest_group >= 4
	case 10:
		isValid = isValid && dice.largest_group == 1 && dice.Counts[5] == 0
	case 11:
		isValid = isValid && dice.largest_group == 1 && dice.Counts[0] == 0
	case 12:
		isValid = isValid && dice.group_count == 2 && dice.largest_group == 3
	case 14:
		isValid = isValid && dice.largest_group == 5
	default:
		return isValid
	}

	return isValid
}

func ValidCategories(sheet Sheet, dice Dice) []int {
	valid := []int{}
	for i := range sheet.Categories {
		if isValid(sheet.Categories[i], dice, i) {
			valid = append(valid, i)
		}
	}
	return valid
}

func ActiveCategories(sheet Sheet) []int {
	active := []int{}
	for i := range sheet.Categories {
		if !sheet.Categories[i].Used {
			active = append(active, i)
		}
	}
	return active
}

func categoryPoints(dice Dice, index int) int {
	points := 0
	switch index {
	// ones to sixes
	case 0, 1, 2, 3, 4, 5:
		points = dice.sums[index]
	// pairs
	case 6:
		// go backwards to use the highest face available
		for i := 5; i >= 0; i-- {
			if dice.Counts[i] >= 2 {
				points = 2 * (i + 1)
				break
			}
		}
	// two pairs
	case 7:
		for i := range dice.Counts {
			if dice.Counts[i] >= 2 {
				points = 2 * (i + 1)
			}
		}
	// three of a kind
	case 8:
		points = 3 * (dice.largest_group_index + 1)
	// four of a kind
	case 9:
		points = 4 * (dice.largest_group_index + 1)
	// all categories that use all 5 dice
	case 10, 11, 12, 13:
		points = dice.total_sum
	case 14:
		points = 50
	}
	return points
}

func ConsumeCategory(sheet Sheet, dice Dice, index int) Sheet {
	if sheet.Categories[index].Used == true {
		options := ActiveCategories(sheet)
		i := rand.Intn(len(options))
		index = options[i]
	}
	isValid := isValid(sheet.Categories[index], dice, index)
	if isValid {
		points := categoryPoints(dice, index)
		sheet.Categories[index].Points = points
		sheet.total_sum += points
		if index <= 5 {
			sheet.upper_section_sum += points
		}
	}
	sheet.Categories[index].Used = true
	sheet.Categories_left -= 1
	return sheet
}

func GameScore(sheet Sheet) int {
	score := sheet.total_sum
	if sheet.upper_section_sum >= 63 {
		score += 50
	}
	return score
}
