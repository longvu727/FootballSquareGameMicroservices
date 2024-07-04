package routes

import (
	"bytes"
	"footballsquaregamemicroservices/app"
	mockfootballsquaregameapp "footballsquaregamemicroservices/app/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	footballsquaregamemicroservices "github.com/longvu727/FootballSquaresLibs/services/football_square_game_microservices"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RoutesTestSuite struct {
	suite.Suite
}

func (suite *RoutesTestSuite) TestCreateFootballSquareGame() {

	url := "/CreateFootballSquareGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"game_id": 123, "square_id": 345, "square_size": 10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	mockFootballSquareGame := mockfootballsquaregameapp.NewMockFootballSquareGame(ctrl)
	mockFootballSquareGame.EXPECT().
		CreateDBFootballSquareGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&app.CreateFootballSquareGameResponse{FootballSquaresGameIDs: []int64{1, 2, 3, 4}}, nil)

	routes := Routes{Apps: mockFootballSquareGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetFootballSquareGame() {

	url := "/GetFootballSquareGame"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"football_square_game_id":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnFootballSquareGame := &app.GetFootballSquareGameResponse{}
	returnFootballSquareGame.FootballSquaresGameID = 10
	returnFootballSquareGame.ColumnIndex = 1
	returnFootballSquareGame.RowIndex = 1
	returnFootballSquareGame.WinnerQuaterNumber = 0
	returnFootballSquareGame.Winner = false
	returnFootballSquareGame.UserID = 123
	returnFootballSquareGame.SquareID = 321
	returnFootballSquareGame.GameID = 234
	returnFootballSquareGame.UserName = "LongUserName"
	returnFootballSquareGame.UserAlias = "usr"

	mockFootballSquareGame := mockfootballsquaregameapp.NewMockFootballSquareGame(ctrl)
	mockFootballSquareGame.EXPECT().
		GetFootballSquareGame(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnFootballSquareGame, nil)

	routes := Routes{Apps: mockFootballSquareGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestGetFootballSquareGameByGameID() {

	url := "/GetFootballSquareGameByGameID"
	ctrl := gomock.NewController(suite.T())

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(`{"game_id":10}`)))
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	returnFootballSquareGames := &app.GetFootballSquareGamesResponse{
		FootballSquareGames: []footballsquaregamemicroservices.FootballSquareGameElement{
			{
				FootballSquaresGameID: 10,
				ColumnIndex:           1,
				RowIndex:              1,
				WinnerQuaterNumber:    0,
				Winner:                false,
				UserID:                123,
				SquareID:              321,
				GameID:                234,
				UserName:              "LongUserName",
				UserAlias:             "usr",
			},
			{
				FootballSquaresGameID: 10,
				ColumnIndex:           1,
				RowIndex:              1,
				WinnerQuaterNumber:    0,
				Winner:                false,
				UserID:                123,
				SquareID:              321,
				GameID:                234,
				UserName:              "LongUserName",
				UserAlias:             "usr",
			},
		},
	}

	mockFootballSquareGame := mockfootballsquaregameapp.NewMockFootballSquareGame(ctrl)
	mockFootballSquareGame.EXPECT().
		GetFootballSquareGameByGameID(gomock.Any(), gomock.Any()).
		Times(1).
		Return(returnFootballSquareGames, nil)

	routes := Routes{Apps: mockFootballSquareGame}
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(http.MethodPost+" "+url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func (suite *RoutesTestSuite) TestHome() {

	url := "/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	suite.NoError(err)

	httpRecorder := httptest.NewRecorder()

	routes := NewRoutes()
	serveMux := routes.Register(nil)

	handler, pattern := serveMux.Handler(req)
	suite.Equal(url, pattern)

	handler.ServeHTTP(httpRecorder, req)

	suite.Equal(httpRecorder.Code, http.StatusOK)
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}
