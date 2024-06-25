// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type FootballSquareGame struct {
	FootballSquareGameID int32
	GameID               sql.NullInt32
	SquareID             sql.NullInt32
	UserID               sql.NullInt32
	Winner               sql.NullBool
	WinnerQuarterNumber  sql.NullInt16
	RowIndex             sql.NullInt32
	ColumnIndex          sql.NullInt32
	Created              sql.NullTime
	Updated              sql.NullTime
}

type Game struct {
	GameID   int32
	GameGuid string
	Sport    sql.NullString
	TeamA    sql.NullString
	TeamB    sql.NullString
	Created  sql.NullTime
	Updated  sql.NullTime
}

type Square struct {
	SquareID     int32
	SquareGuid   string
	SquareSize   sql.NullInt32
	RowPoints    sql.NullString
	ColumnPoints sql.NullString
	Created      sql.NullTime
	Updated      sql.NullTime
}

type User struct {
	UserID     int32
	UserGuid   string
	Ip         sql.NullString
	DeviceName sql.NullString
	UserName   sql.NullString
	Alias      sql.NullString
	Created    sql.NullTime
	Updated    sql.NullTime
}
