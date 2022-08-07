package rps_web

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

var phrases = ResultPhrases{
	PLAYERWINS:   {"You're lucky!", "Great!", "Congrats!", "You should buy a lottery ticket", "Nice work!", "Good job!mv "},
	COMPUTERWINS: {"Too Bad!", "Next time you'll be luckier...", "Try again...", "I'll be back..."},
	DRAW:         {"Gegious minds think the same.", "Nobody won.", "Ok."},
}

type Round struct {
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	ResultPhrase   string `json:"result_phrase"`
}

type ResultPhrases map[int][]string

func PlayRound(playerValue int) Round {
	return PlayRoundConfigurable(playerValue, randomValueGenerator, 3)
}

func PlayRoundConfigurable(playerValue int, compValueGenerator func(int) int, itemsNumber int) Round {
	var result Round
	var winner int

	computerValue := compValueGenerator(itemsNumber)

	computerChoice := CovertComputerChoiceToString(computerValue)
	result.ComputerChoice = fmt.Sprintf("Computer chooses %s", computerChoice)

	playerScore, computerScore := RoundWinner(playerValue, computerValue)

	result.RoundResult, winner = FTotalResult(playerScore, computerScore)

	result.ResultPhrase = getPhrase(winner, phrases)

	return result
}

func getPhrase(winnerConst int, phrases ResultPhrases) string {
	variantsNumber := len(phrases[winnerConst])
	choosedVariant := randomValueGenerator(variantsNumber)
	return phrases[winnerConst][choosedVariant]
}

func randomValueGenerator(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RoundWinner(player1, player2 int) (resPlayer1 int, resPlayer2 int) {
	if player1 == player2 {
		return 0, 0
	} else if (player1+1)%3 == player2 {
		return 0, 1
	} else {
		return 1, 0
	}
}

func CovertComputerChoiceToString(computerChoiceNum int) string {
	var computerChoice string
	switch computerChoiceNum {
	case 0:
		computerChoice = "ROCK"
		break
	case 1:
		computerChoice = "PAPER"
		break
	case 2:
		computerChoice = "SCISSORS"
		break
	}
	return computerChoice
}

func FTotalResult(playerScore, computerScore int) (string, int) {
	if playerScore > computerScore {
		return "Human wins!", PLAYERWINS
	} else if playerScore == computerScore {
		return "It's draw.", DRAW
	} else {
		return "Computer wins!", COMPUTERWINS
	}
}
