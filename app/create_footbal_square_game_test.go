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

type CreateFootballSquareGameTestSuite struct {
	suite.Suite
}

func TestCreateFootballSquareGameTestSuite(t *testing.T) {
	suite.Run(t, new(CreateFootballSquareGameTestSuite))
}

func (suite *CreateFootballSquareGameTestSuite) TestCreateFootballSquareGame() {
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateFootballSquareGame(gomock.Any(), gomock.Any()).
		AnyTimes().
		Return(rand.Int63n(1000), nil)

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	squareSize := getSquareSizeNonZero()
	createSquareParams := CreateFootballSquareGameParams{
		GameID:     rand.Int31n(1000),
		SquareID:   rand.Int31n(1000),
		SquareSize: int32(squareSize),
	}

	footballSquareGame, err := NewFootballSquareGameApp().CreateDBFootballSquareGame(createSquareParams, resources)
	suite.NoError(err)

	suite.Equal(len(footballSquareGame.FootballSquaresGameIDs), squareSize*squareSize)
	suite.Greater(len(footballSquareGame.ToJson()), 0)
}

func (suite *CreateFootballSquareGameTestSuite) TestCreateFootballSquareGameDBError() {

	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()

	mockMySQL := mockdb.NewMockMySQL(ctrl)

	mockMySQL.EXPECT().
		CreateFootballSquareGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(int64(0), errors.New("test error"))

	config, err := util.LoadConfig("../env", "app", "env")
	suite.NoError(err)

	resources := resources.NewResources(config, mockMySQL, services.NewServices(), context.Background())

	createSquareParams := CreateFootballSquareGameParams{
		GameID:     rand.Int31n(1000),
		SquareID:   rand.Int31n(1000),
		SquareSize: int32(getSquareSizeNonZero()),
	}

	_, err = NewFootballSquareGameApp().CreateDBFootballSquareGame(createSquareParams, resources)
	suite.Error(err)
}

func getSquareSizeNonZero() int {
	return rand.Intn(9) + 1
}
