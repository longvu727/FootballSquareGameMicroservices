package app

import (
	"database/sql"
	"encoding/json"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type ReserveFootballSquareParams struct {
	UserID               int32 `json:"user_id"`
	FootballSquareGameID int32 `json:"football_square_game_id"`
	RowIndex             int32 `json:"row_index"`
	ColumnIndex          int32 `json:"column_index"`
}

type ReserveFootballSquareResponse struct {
	Reserved     bool   `json:"reserved"`
	ErrorMessage string `json:"error_message"`
}

func (response ReserveFootballSquareResponse) ToJson() []byte {
	jsonStr, _ := json.Marshal(response)
	return jsonStr
}

func (footballSquareGameApp *FootballSquareGameApp) ReserveFootballSquare(reserveFootballSquareParams ReserveFootballSquareParams, resources *resources.Resources) (*ReserveFootballSquareResponse, error) {
	var reserveFootballSquareGameResponse ReserveFootballSquareResponse

	err := resources.DB.ReserveFootballSquareByGameIDRowIndexColumnIndex(
		resources.Context,
		db.ReserveFootballSquareByGameIDRowIndexColumnIndexParams{
			UserID:               sql.NullInt32{Int32: reserveFootballSquareParams.UserID, Valid: true},
			FootballSquareGameID: reserveFootballSquareParams.FootballSquareGameID,
			RowIndex:             sql.NullInt32{Int32: reserveFootballSquareParams.RowIndex, Valid: true},
			ColumnIndex:          sql.NullInt32{Int32: reserveFootballSquareParams.ColumnIndex, Valid: true},
		},
	)
	if err != nil {
		return &reserveFootballSquareGameResponse, err
	}

	reserveFootballSquareGameResponse.Reserved = true

	return &reserveFootballSquareGameResponse, nil
}
