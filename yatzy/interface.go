package yatzy

// outputs 15 + 18 = 33 elements
func InputVec(game Yatzy) []float64 {
	vec := []float64{}
	for i := range game.sheet.Categories {
		category := game.sheet.Categories[i]
		if !category.Used {
			vec = append(vec, 1.0)
		} else {
			vec = append(vec, 0.0)
		}
	}
	for i := range game.dice.Counts {
		count := game.dice.Counts[i]

		if count&1 != 0 {
			vec = append(vec, 1.0)
		} else {
			vec = append(vec, 0.0)
		}

		if count&2 != 0 {
			vec = append(vec, 1.0)
		} else {
			vec = append(vec, 0.0)
		}

		if count&4 != 0 {
			vec = append(vec, 1.0)
		} else {
			vec = append(vec, 0.0)
		}
	}
	return vec
}
