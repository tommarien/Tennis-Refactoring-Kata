package tenniskata

import (
	"fmt"
	"math"
	"strconv"
)

type tennisGame1 struct {
	player1 Player
	player2 Player
}

type Player struct {
	name  string
	score int
}

func (player *Player) ScoreName() string {
	switch player.score {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"

	default:
		return strconv.Itoa(player.score)
	}
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1: Player{
			name: player1Name,
		},
		player2: Player{
			name: player2Name,
		},
	}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	switch playerName {
	case game.player1.name:
		game.player1.score += 1
	case game.player2.name:
		game.player2.score += 1
	}
}

func (game *tennisGame1) GetScore() string {
	switch {
	case game.player1.score == game.player2.score:
		switch {
		case game.player1.score < 3:
			return fmt.Sprintf("%s-All", game.player1.ScoreName())
		default:
			return "Deuce"
		}
	case game.player1.score >= 4 || game.player2.score >= 4:
		scoreDifference := game.player1.score - game.player2.score
		var leadingPlayer *Player

		if scoreDifference > 0 {
			leadingPlayer = &game.player1
		} else {
			leadingPlayer = &game.player2
		}

		switch math.Abs(float64(scoreDifference)) {
		case 1:
			return fmt.Sprintf("Advantage %s", leadingPlayer.name)
		default:
			return fmt.Sprintf("Win for %s", leadingPlayer.name)
		}

	default:
		return fmt.Sprintf("%s-%s", game.player1.ScoreName(), game.player2.ScoreName())
	}
}
