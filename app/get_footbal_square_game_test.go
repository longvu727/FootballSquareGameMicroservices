package app

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"testing"

	"math/rand"

	"github.com/golang/mock/gomock"
	"github.com/longvu727/FootballSquaresLibs/DB/db"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type GetFootballSquareGameTestSuite struct {
	suite.Suite
}

func TestGetFootballSquareGameTestSuite(t *testing.T) {
	suite.Run(t, new(GetFootballSquareGameTestSuite))
}

func (suite *GetFootballSquareGameTestSuite) TestGetFootballSquareGame() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	randomFootballSquareGame := randomFootballSquareGame()

	mockMySQL.EXPECT().
		GetFootballSquareGame(gomock.Any(), gomock.Any()).
		AnyTimes().
		Return(randomFootballSquareGame, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	getSquareParams := GetFootballSquareGameParams{FootballSquaresGameID: randomFootballSquareGame.FootballSquareGameID}

	footballSquareGame, err := NewFootballSquareGameApp().GetFootballSquareGame(getSquareParams, resources)
	suite.NoError(err)

	suite.Equal(randomFootballSquareGame.GameID.Int32, int32(footballSquareGame.GameID))
	suite.Equal(randomFootballSquareGame.FootballSquareGameID, int32(footballSquareGame.FootballSquareGameID))
	suite.Equal(randomFootballSquareGame.SquareID.Int32, int32(footballSquareGame.SquareID))
	suite.Equal(randomFootballSquareGame.UserID.Int32, int32(footballSquareGame.UserID))
	suite.Equal(randomFootballSquareGame.Winner.Bool, footballSquareGame.Winner)
	suite.Equal(randomFootballSquareGame.WinnerQuarterNumber.Int16, int16(footballSquareGame.WinnerQuarterNumber))
	suite.Equal(randomFootballSquareGame.RowIndex.Int32, int32(footballSquareGame.RowIndex))
	suite.Equal(randomFootballSquareGame.ColumnIndex.Int32, int32(footballSquareGame.ColumnIndex))

	suite.Greater(len(footballSquareGame.ToJson()), 0)
}

func (suite *GetFootballSquareGameTestSuite) TestGetFootballSquareGameDBError() {

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetFootballSquareGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(db.GetFootballSquareGameRow{}, errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	getSquareParams := GetFootballSquareGameParams{}

	_, err = NewFootballSquareGameApp().GetFootballSquareGame(getSquareParams, resources)
	suite.Error(err)
}

func (suite *GetFootballSquareGameTestSuite) TestGetFootballSquareGameByGameID() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	randomFootballSquareGames := randomFootballSquareGames()

	mockMySQL.EXPECT().
		GetFootballSquareGameByGameID(gomock.Any(), gomock.Any()).
		AnyTimes().
		Return(randomFootballSquareGames, nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	getSquareParams := GetFootballSquareGameByGameIDParams{GameID: randomFootballSquareGames[0].GameID.Int32}

	footballSquareGames, err := NewFootballSquareGameApp().GetFootballSquareGameByGameID(getSquareParams, resources)
	suite.NoError(err)

	suite.Equal(len(randomFootballSquareGames), len(footballSquareGames.FootballSquareGames))

	suite.Greater(len(footballSquareGames.ToJson()), 0)
}

func (suite *GetFootballSquareGameTestSuite) TestGetFootballSquareGameByGameIDDBError() {

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		GetFootballSquareGameByGameID(gomock.Any(), gomock.Any()).
		Times(1).
		Return([]db.GetFootballSquareGameByGameIDRow{}, errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	getSquareParams := GetFootballSquareGameByGameIDParams{}

	_, err = NewFootballSquareGameApp().GetFootballSquareGameByGameID(getSquareParams, resources)
	suite.Error(err)
}

func randomFootballSquareGame() db.GetFootballSquareGameRow {
	return db.GetFootballSquareGameRow{
		GameID:               sql.NullInt32{Valid: true, Int32: rand.Int31n(1000)},
		FootballSquareGameID: rand.Int31n(1000),
		SquareID:             sql.NullInt32{Valid: true, Int32: rand.Int31n(1000)},
		UserID:               sql.NullInt32{Valid: true, Int32: rand.Int31n(1000)},
		Winner:               sql.NullBool{Valid: true, Bool: false},
		WinnerQuarterNumber:  sql.NullInt16{},
		RowIndex:             sql.NullInt32{Valid: true, Int32: rand.Int31n(10)},
		ColumnIndex:          sql.NullInt32{Valid: true, Int32: rand.Int31n(10)},
	}
}

func randomFootballSquareGames() []db.GetFootballSquareGameByGameIDRow {
	getFootballSquareGameRow := []db.GetFootballSquareGameByGameIDRow{}

	squareSize := rand.Intn(99) + 1
	for i := 0; i < squareSize; i++ {
		randomFootballSquareGame := randomFootballSquareGame()
		getFootballSquareGameRow = append(getFootballSquareGameRow, db.GetFootballSquareGameByGameIDRow{
			GameID:               randomFootballSquareGame.GameID,
			FootballSquareGameID: randomFootballSquareGame.FootballSquareGameID,
			SquareID:             randomFootballSquareGame.SquareID,
			UserID:               randomFootballSquareGame.UserID,
			Winner:               randomFootballSquareGame.Winner,
			WinnerQuarterNumber:  randomFootballSquareGame.WinnerQuarterNumber,
			RowIndex:             randomFootballSquareGame.RowIndex,
			ColumnIndex:          randomFootballSquareGame.ColumnIndex,

			UserName: sql.NullString{
				String: "user" + strconv.Itoa(rand.Intn(1000)), Valid: true,
			},
			Alias: sql.NullString{
				String: "u" + strconv.Itoa(rand.Intn(1000)), Valid: true,
			},
		})
	}

	return getFootballSquareGameRow
}
