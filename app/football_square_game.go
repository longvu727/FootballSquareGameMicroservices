package app

import "github.com/longvu727/FootballSquaresLibs/util/resources"

type FootballSquareGame interface {
	GetFootballSquareGame(getFootballSquareGameParams GetFootballSquareGameParams, resources *resources.Resources) (*GetFootballSquareGameResponse, error)
	GetFootballSquareGameByGameID(getFootballSquareGameByGameIDParams GetFootballSquareGameByGameIDParams, resources *resources.Resources) (*GetFootballSquareGamesResponse, error)
	CreateDBFootballSquareGame(createFootballSquareGameParams CreateFootballSquareGameParams, resources *resources.Resources) (*CreateFootballSquareGameResponse, error)
}

type FootballSquareGameApp struct{}

func NewFootballSquareGameApp() FootballSquareGame {
	return &FootballSquareGameApp{}
}
