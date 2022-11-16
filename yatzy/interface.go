package yatzy

const outputSize = 15

// outputs 15 + 18 = 33 elements
func InputVec(dice Dice, sheet Sheet) []float64 {
	vec := []float64{}
	for i := range sheet.Categories {
		category := sheet.Categories[i]
		if !category.Used {
			vec = append(vec, 1.0)
		} else {
			vec = append(vec, 0.0)
		}
	}
	for i := range dice.Counts {
		count := dice.Counts[i]

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

func desiredRerollOutput(output []float64) []float64 {
	desired := []float64{}
	for i := range output {
		if i < 5 {
			if output[i] >= 0.5 {
				desired = append(desired, 1.0)
			} else {
				desired = append(desired, 0.0)
			}
		} else {
			desired = append(desired, output[i])
		}
	}
	return desired
}

/*
func DesiredCategoryOutput(output *mat.VecDense, correct bool, index int) *mat.VecDense {
	desired := []float64{}
	for i := 0; i < outputSize; i += 1 {
		if i < 5 {
			desired = append(desired, output.AtVec(i))
		} else {
			if correct {
				if i == index {
					desired = append(desired, 1.0)
				} else {
					desired = append(desired, 0.0)
				}
			} else {
				if i == index {
					desired = append(desired, 0.0)
				} else {
					desired = append(desired, output.AtVec(i))
				}
			}
		}
	}
	return mat.NewVecDense(len(desired), desired)
}
*/
