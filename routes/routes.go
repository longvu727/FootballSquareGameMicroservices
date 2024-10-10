package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"footballsquaregamemicroservices/app"
	"log"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type RoutesInterface interface {
	Register(resources *resources.Resources) *http.ServeMux
}

type Routes struct {
	Apps app.FootballSquareGame
}

type Handler = func(writer http.ResponseWriter, request *http.Request, resources *resources.Resources)

func NewRoutes() RoutesInterface {
	return &Routes{
		Apps: app.NewFootballSquareGameApp(),
	}
}

func (routes *Routes) Register(resources *resources.Resources) *http.ServeMux {
	log.Println("Registering routes")
	mux := http.NewServeMux()

	routesHandlersMap := map[string]Handler{
		"/": routes.home,

		http.MethodPost + " /CreateFootballSquareGame":      routes.createFootballSquareGame,
		http.MethodPost + " /GetFootballSquareGame":         routes.getFootballSquareGame,
		http.MethodPost + " /GetFootballSquareGameByGameID": routes.getFootballSquareGameByGameID,
		http.MethodPost + " /ReserveFootballSquare":         routes.reserveFootballSquare,
	}

	for route, handler := range routesHandlersMap {
		mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, resources)
		})
	}

	return mux
}

func (routes *Routes) home(writer http.ResponseWriter, _ *http.Request, resources *resources.Resources) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func (routes *Routes) createFootballSquareGame(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var createFootballSquareGameParams app.CreateFootballSquareGameParams
	json.NewDecoder(request.Body).Decode(&createFootballSquareGameParams)

	createSquareResponse, err := routes.Apps.CreateDBFootballSquareGame(createFootballSquareGameParams, resources)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.ErrorMessage = `Unable to create FootballSquareGame`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func (routes *Routes) getFootballSquareGame(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	var getFootballSquareGameParams app.GetFootballSquareGameParams
	json.NewDecoder(request.Body).Decode(&getFootballSquareGameParams)

	getFootballSquareGameResponse, err := routes.Apps.GetFootballSquareGame(getFootballSquareGameParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getFootballSquareGameResponse.ErrorMessage = `FootballSquareGame not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getFootballSquareGameResponse.ErrorMessage = `Unable to get FootballSquareGame`
		writer.Write(getFootballSquareGameResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getFootballSquareGameResponse.ToJson())
}

func (routes *Routes) getFootballSquareGameByGameID(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")
	var getFootballSquareGameByGameIDParams app.GetFootballSquareGameByGameIDParams
	json.NewDecoder(request.Body).Decode(&getFootballSquareGameByGameIDParams)

	getFootballSquareResponse, err := routes.Apps.GetFootballSquareGameByGameID(getFootballSquareGameByGameIDParams, resources)

	if err != nil && err == sql.ErrNoRows {
		getFootballSquareResponse.ErrorMessage = `FootballSquareGame not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getFootballSquareResponse.ErrorMessage = `Unable to get FootballSquareGame`
		writer.Write(getFootballSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getFootballSquareResponse.ToJson())
}

func (routes *Routes) reserveFootballSquare(writer http.ResponseWriter, request *http.Request, resources *resources.Resources) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")
	var reserveFootballSquareParams app.ReserveFootballSquareParams
	json.NewDecoder(request.Body).Decode(&reserveFootballSquareParams)

	reserveFootballSquareResponse, err := routes.Apps.ReserveFootballSquare(reserveFootballSquareParams, resources)

	if err != nil && err == sql.ErrNoRows {
		reserveFootballSquareResponse.ErrorMessage = `guid not found`
	} else if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		reserveFootballSquareResponse.ErrorMessage = `Unable to reserve FootballSquareGame`
		writer.Write(reserveFootballSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(reserveFootballSquareResponse.ToJson())
}
