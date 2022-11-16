package nn

/*
func playGame(network *Network) int {
	sheet := yatzy.NewSheet()

	for i := 0; i < 10; i += 1 {
		// roll 1
		dice := yatzy.Roll()
		input := yatzy.InputVec(dice, sheet)
		ForwardProp(network, input)

		// roll 2
		dice = yatzy.Reroll(dice, OutputRerolls(network.A3))
		input = yatzy.InputVec(dice, sheet)
		ForwardProp(network, input)

		// roll 3
		dice = yatzy.Reroll(dice, OutputRerolls(network.A3))
		input = yatzy.InputVec(dice, sheet)
		ForwardProp(network, input)

		// select category
		sheet = yatzy.ConsumeCategory(sheet, dice, OutputCategory(network.A3, input))
	}

	return yatzy.GameScore(sheet)
}

func NGameAverage(network *Network, n int) int {
	total := 0
	for i := 0; i < n; i += 1 {
		total += playGame(network)
	}
	return total / n
}
*/
