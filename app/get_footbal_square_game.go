package app

import (
	"database/sql"
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type GetFootballSquareGameParams struct {
	FootballSquaresGameID int32 `json:"football_square_game_id"`
}

type GetFootballSquareGamesResponse struct {
	FootballSquareGames []services.FootballSquareGameElement `json:"football_squares"`
	ErrorMessage        string                               `json:"error_message"`
}

type GetFootballSquareGameByGameIDParams struct {
	GameID int32 `json:"game_id"`
}

type GetFootballSquareGameResponse struct {
	services.FootballSquareGameElement
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

func (footballSquareGameApp *FootballSquareGameApp) GetFootballSquareGame(getFootballSquareGameParams GetFootballSquareGameParams, resources *resources.Resources) (*GetFootballSquareGameResponse, error) {
	var getFootballSquareGameResponse GetFootballSquareGameResponse

	footballGameRow, err := resources.DB.GetFootballSquareGame(resources.Context, getFootballSquareGameParams.FootballSquaresGameID)
	if err != nil {
		return &getFootballSquareGameResponse, err
	}

	getFootballSquareGameResponse.FootballSquareGameID = int(footballGameRow.FootballSquareGameID)
	getFootballSquareGameResponse.ColumnIndex = int(footballGameRow.ColumnIndex.Int32)
	getFootballSquareGameResponse.RowIndex = int(footballGameRow.RowIndex.Int32)
	getFootballSquareGameResponse.WinnerQuarterNumber = int(footballGameRow.WinnerQuarterNumber.Int16)
	getFootballSquareGameResponse.Winner = footballGameRow.Winner.Bool
	getFootballSquareGameResponse.UserID = int(footballGameRow.UserID.Int32)
	getFootballSquareGameResponse.SquareID = int(footballGameRow.SquareID.Int32)
	getFootballSquareGameResponse.GameID = int(footballGameRow.GameID.Int32)

	return &getFootballSquareGameResponse, nil
}

func (footballSquareGameApp *FootballSquareGameApp) GetFootballSquareGameByGameID(getFootballSquareGameByGameIDParams GetFootballSquareGameByGameIDParams, resources *resources.Resources) (*GetFootballSquareGamesResponse, error) {
	var getFootballSquareGamesResponse GetFootballSquareGamesResponse

	footballGameRows, err := resources.DB.GetFootballSquareGameByGameID(
		resources.Context,
		sql.NullInt32{
			Int32: getFootballSquareGameByGameIDParams.GameID,
			Valid: true,
		},
	)
	if err != nil {
		return &getFootballSquareGamesResponse, err
	}

	for _, footballGameRow := range footballGameRows {
		getFootballSquareGameElement := services.FootballSquareGameElement{
			FootballSquareGameID: int(footballGameRow.FootballSquareGameID),
			ColumnIndex:          int(footballGameRow.ColumnIndex.Int32),
			RowIndex:             int(footballGameRow.RowIndex.Int32),
			WinnerQuarterNumber:   int(footballGameRow.WinnerQuarterNumber.Int16),
			Winner:               footballGameRow.Winner.Bool,
			UserID:               int(footballGameRow.UserID.Int32),
			SquareID:             int(footballGameRow.SquareID.Int32),
			GameID:               int(footballGameRow.GameID.Int32),
			UserName:             footballGameRow.UserName.String,
			UserAlias:            footballGameRow.Alias.String,
		}
		getFootballSquareGamesResponse.FootballSquareGames =
			append(getFootballSquareGamesResponse.FootballSquareGames, getFootballSquareGameElement)
	}
	return &getFootballSquareGamesResponse, nil
}
