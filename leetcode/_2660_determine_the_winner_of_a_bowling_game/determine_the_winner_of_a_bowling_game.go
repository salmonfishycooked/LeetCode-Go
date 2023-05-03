package _2660_determine_the_winner_of_a_bowling_game

func isWinner(player1 []int, player2 []int) int {
	score1, score2 := cal(player1), cal(player2)
	if score1 > score2 {
		return 1
	} else if score1 < score2 {
		return 2
	}
	return 0
}

func cal(arr []int) (score int) {
	for i, v := range arr {
		if (i > 0 && arr[i-1] == 10) || (i > 1 && arr[i-2] == 10) {
			score += 2 * v
			continue
		}
		score += v
	}
	return
}
