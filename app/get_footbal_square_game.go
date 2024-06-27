package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	footballsquaregameservices "github.com/longvu727/FootballSquaresLibs/services/football_square_game_microservices"
)

type GetFootballSquareGameParams struct {
	FootballSquaresGameID int32 `json:"football_square_game_id"`
}

type GetFootballSquareGamesResponse struct {
	FootballSquareGames footballsquaregameservices.FootballSquares `json:"football_square"`
	ErrorMessage        string                                     `json:"error_message"`
}

type GetFootballSquareGameByGameIDParams struct {
	GameID int32 `json:"game_id"`
}

type GetFootballSquareGameResponse struct {
	footballsquaregameservices.FootballSquareGameElement
	ErrorMessage string `json:"error_message"`
}

func (response GetFootballSquareGameResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (response GetFootballSquareGamesResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func GetFootballSquareGame(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*GetFootballSquareGameResponse, error) {
	var getFootballSquareGameResponse GetFootballSquareGameResponse
	var getFootballSquareGameParams GetFootballSquareGameParams
	json.NewDecoder(request.Body).Decode(&getFootballSquareGameParams)

	footballGameRow, err := dbConnect.QUERIES.GetFootballSquareGame(ctx, getFootballSquareGameParams.FootballSquaresGameID)
	if err != nil {
		return &getFootballSquareGameResponse, err
	}

	getFootballSquareGameResponse.FootballSquaresGameID = int(footballGameRow.FootballSquareGameID)
	getFootballSquareGameResponse.ColumnIndex = int(footballGameRow.ColumnIndex.Int32)
	getFootballSquareGameResponse.RowIndex = int(footballGameRow.RowIndex.Int32)
	getFootballSquareGameResponse.WinnerQuaterNumber = int(footballGameRow.WinnerQuarterNumber.Int16)
	getFootballSquareGameResponse.Winner = footballGameRow.Winner.Bool
	getFootballSquareGameResponse.UserID = int(footballGameRow.UserID.Int32)
	getFootballSquareGameResponse.SquareID = int(footballGameRow.SquareID.Int32)
	getFootballSquareGameResponse.GameID = int(footballGameRow.GameID.Int32)

	return &getFootballSquareGameResponse, nil
}

func GetFootballSquareGameByGameID(ctx context.Context, request *http.Request, dbConnect *db.MySQL) (*GetFootballSquareGamesResponse, error) {
	var getFootballSquareGamesResponse GetFootballSquareGamesResponse
	var getFootballSquareGameByGameIDParams GetFootballSquareGameByGameIDParams
	json.NewDecoder(request.Body).Decode(&getFootballSquareGameByGameIDParams)

	footballGameRows, err := dbConnect.QUERIES.GetFootballSquareGameByGameID(
		ctx,
		sql.NullInt32{
			Int32: getFootballSquareGameByGameIDParams.GameID,
			Valid: true,
		},
	)
	if err != nil {
		return &getFootballSquareGamesResponse, err
	}

	for _, footballGameRow := range footballGameRows {
		getFootballSquareGameElement := footballsquaregameservices.FootballSquareGameElement{
			FootballSquaresGameID: int(footballGameRow.FootballSquareGameID),
			ColumnIndex:           int(footballGameRow.ColumnIndex.Int32),
			RowIndex:              int(footballGameRow.RowIndex.Int32),
			WinnerQuaterNumber:    int(footballGameRow.WinnerQuarterNumber.Int16),
			Winner:                footballGameRow.Winner.Bool,
			UserID:                int(footballGameRow.UserID.Int32),
			SquareID:              int(footballGameRow.SquareID.Int32),
			GameID:                int(footballGameRow.GameID.Int32),
		}
		getFootballSquareGamesResponse.FootballSquareGames =
			append(getFootballSquareGamesResponse.FootballSquareGames, getFootballSquareGameElement)
	}
	return &getFootballSquareGamesResponse, nil
}
