package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type Message struct {
	Status  int    `json:status`
	Message string `json:message`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/").
		Name("Index").
		HandlerFunc(Index)

	router.
		Methods("POST").
		Path("/data").
		Name("Data").
		HandlerFunc(Data)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Status:  http.StatusOK,
		Message: "OK",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

func Data(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	requestDataJson, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Marshal error: %v\n", err)
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO damdata (info) VALUES ($1)", requestDataJson)
	if err != nil {
		log.Fatalf("An error occurred: %v\n", err)
	}

	log.Printf("%+v", requestData)
}
