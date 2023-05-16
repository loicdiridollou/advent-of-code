package main

type State struct {
	position1, position2, score1, score2, player_turn int
}

var dicesum_odds [][2]int = [][2]int{
	{3, 1},
	{4, 3},
	{5, 6},
	{6, 7},
	{7, 6},
	{8, 3},
	{9, 1},
}

func newPosition(position int, dice_sum int) int {
	return (position+dice_sum-1)%10 + 1
}

func countWins(s State) []int {
	if s.score1 >= 21 {
		return []int{1, 0}
	}
	if s.score2 >= 21 {
		return []int{0, 1}
	}

	wins := []int{0, 0}
	var new_s State
	for _, odd := range dicesum_odds {
		if s.player_turn == 1 {
			pos := newPosition(s.position1, odd[0])
			new_s = State{pos, s.position2, s.score1 + pos, s.score2, 2}
		} else {
			pos := newPosition(s.position2, odd[0])
			new_s = State{s.position1, pos, s.score1, s.score2 + pos, 1}
		}
		tmp_wins := countWins(new_s)
		wins[0] += tmp_wins[0] * odd[1]
		wins[1] += tmp_wins[1] * odd[1]
	}

	return wins
}

func MaxInt(a []int) int {
	if a[0] > a[1] {
		return a[0]
	}
	return a[1]
}

func part2(input string) int {
	positions := prepareData(input)
	return MaxInt(countWins(State{positions[0], positions[1], 0, 0, 1}))
}
