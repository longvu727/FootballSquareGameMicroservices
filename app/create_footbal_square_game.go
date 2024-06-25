package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
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

func CreateDBFootballSquareGame(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*CreateFootballSquareGameResponse, error) {
	var createFootballSquareGameParams CreateFootballSquareGameParams
	json.NewDecoder(request.Body).Decode(&createFootballSquareGameParams)

	var createFootballSquareGameResponse CreateFootballSquareGameResponse

	footballSquareGameIDs, err := generateFootballSquareGame(ctx, createFootballSquareGameParams.SquareSize, dbConnect, createFootballSquareGameParams.GameID, createFootballSquareGameParams.SquareID)
	if err != nil {
		return &createFootballSquareGameResponse, err
	}

	createFootballSquareGameResponse.FootballSquaresGameIDs = footballSquareGameIDs

	return &createFootballSquareGameResponse, nil
}

func generateFootballSquareGame(ctx context.Context, squareSize int32, dbConnect *db.MySQL, gameID int32, squareID int32) ([]int64, error) {
	var footballSquareGameIDs []int64

	for row := 1; row <= int(squareSize); row++ {
		for column := 1; column <= int(squareSize); column++ {
			lastID, err := dbConnect.QUERIES.CreateFootballSquareGame(ctx, db.CreateFootballSquareGameParams{
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
