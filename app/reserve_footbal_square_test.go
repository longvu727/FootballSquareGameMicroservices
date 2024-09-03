package app

import (
	"context"
	"errors"
	"testing"

	"math/rand"

	"github.com/golang/mock/gomock"
	mockdb "github.com/longvu727/FootballSquaresLibs/DB/db/mock"
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
	"github.com/stretchr/testify/suite"
)

type ReserveFootballSquareGameTestSuite struct {
	suite.Suite
}

func TestReserveFootballSquareGameTestSuite(t *testing.T) {
	suite.Run(t, new(ReserveFootballSquareGameTestSuite))
}

func (suite *ReserveFootballSquareGameTestSuite) TestReserveFootballSquareGame() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		ReserveFootballSquareByGameIDRowIndexColumnIndex(gomock.Any(), gomock.Any()).
		AnyTimes().
		Return(nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	reserveSquareParams := ReserveFootballSquareParams{
		GameID:      rand.Int31n(1000),
		UserID:      rand.Int31n(1000),
		RowIndex:    rand.Int31n(10),
		ColumnIndex: rand.Int31n(10),
	}

	reserveSquareResponse, err := NewFootballSquareGameApp().ReserveFootballSquare(reserveSquareParams, resources)
	suite.NoError(err)

	suite.True(reserveSquareResponse.Reserved)
	suite.Greater(len(reserveSquareResponse.ToJson()), 0)
}

func (suite *ReserveFootballSquareGameTestSuite) TestReserveFootballSquareGameDBError() {

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		ReserveFootballSquareByGameIDRowIndexColumnIndex(gomock.Any(), gomock.Any()).
		Times(1).
		Return(errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	_, err = NewFootballSquareGameApp().ReserveFootballSquare(ReserveFootballSquareParams{}, resources)
	suite.Error(err)
}
