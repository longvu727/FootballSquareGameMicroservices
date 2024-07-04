package app

import (
	"database/sql"
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type CreateFootballSquareGameParams struct {
	GameID     int32 `json:"game_id"`
	SquareID   int32 `json:"square_id"`
	SquareSize int32 `json:"square_size"`
}
type CreateFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

func (response CreateFootballSquareGameResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (footballSquareGameApp *FootballSquareGameApp) CreateDBFootballSquareGame(createFootballSquareGameParams CreateFootballSquareGameParams, resources *resources.Resources) (*CreateFootballSquareGameResponse, error) {
	var createFootballSquareGameResponse CreateFootballSquareGameResponse

	footballSquareGameIDs, err := footballSquareGameApp.generateFootballSquareGame(
		resources,
		createFootballSquareGameParams.SquareSize,
		createFootballSquareGameParams.GameID,
		createFootballSquareGameParams.SquareID)
	if err != nil {
		return &createFootballSquareGameResponse, err
	}

	createFootballSquareGameResponse.FootballSquaresGameIDs = footballSquareGameIDs

	return &createFootballSquareGameResponse, nil
}

func (footballSquareGameApp *FootballSquareGameApp) generateFootballSquareGame(resources *resources.Resources, squareSize int32, gameID int32, squareID int32) ([]int64, error) {
	var footballSquareGameIDs []int64

	for row := 1; row <= int(squareSize); row++ {
		for column := 1; column <= int(squareSize); column++ {
			lastID, err := resources.DB.CreateFootballSquareGame(resources.Context, db.CreateFootballSquareGameParams{
				GameID:      sql.NullInt32{Int32: gameID, Valid: true},
				SquareID:    sql.NullInt32{Int32: squareID, Valid: true},
				RowIndex:    sql.NullInt32{Int32: int32(row), Valid: true},
				ColumnIndex: sql.NullInt32{Int32: int32(column), Valid: true},
			})

			if err != nil {
				return footballSquareGameIDs, err
			}

			footballSquareGameIDs = append(footballSquareGameIDs, lastID)
		}
	}

	return footballSquareGameIDs, nil
}
