package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"squaremicroservices/app"

	"github.com/longvu727/FootballSquaresLibs/DB/db"
)

type Handler = func(writer http.ResponseWriter, request *http.Request)

func Register(db *db.MySQL, ctx context.Context) {
	log.Println("Registering routes")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		home(w, r)
	})

	http.HandleFunc(http.MethodPost+" /CreateFootballSquareGame", func(w http.ResponseWriter, r *http.Request) {
		createFootballSquareGame(w, r, db, ctx)
	})

	http.HandleFunc(http.MethodPost+" /GetFootballSquareGame", func(w http.ResponseWriter, r *http.Request) {
		getFootballSquareGame(w, r, db, ctx)
	})

	http.HandleFunc(http.MethodPost+" /GetFootballSquareGameByGameID", func(w http.ResponseWriter, r *http.Request) {
		GetFootballSquareGameByGameID(w, r, db, ctx)
	})

}

func home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "{\"Acknowledged\": true}")
}

func createFootballSquareGame(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	createSquareResponse, err := app.CreateDBFootballSquareGame(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		createSquareResponse.ErrorMessage = `Unable to create FootballSquareGame`
		writer.Write(createSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(createSquareResponse.ToJson())
}

func getFootballSquareGame(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	getSquareResponse, err := app.GetFootballSquareGame(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getSquareResponse.ErrorMessage = `Unable to get FootballSquareGame`
		writer.Write(getSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getSquareResponse.ToJson())
}

func GetFootballSquareGameByGameID(writer http.ResponseWriter, request *http.Request, dbConnect *db.MySQL, ctx context.Context) {
	log.Printf("Received request for %s\n", request.URL.Path)

	writer.Header().Set("Content-Type", "application/json")

	getSquareResponse, err := app.GetFootballSquareGameByGameID(ctx, request, dbConnect)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		getSquareResponse.ErrorMessage = `Unable to get FootballSquareGame`
		writer.Write(getSquareResponse.ToJson())
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(getSquareResponse.ToJson())
}
