package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type Message struct {
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

	log.Println("Listening on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Message: "Hello, world!",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

func Data(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now())
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

	message := Message{
		Message: "Data successfully inserted into database",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// log.Printf("Data successfully inserted into database: %+v", requestData)

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

func timeTrack(start time.Time) int64 {
	elapsed := time.Since(start)
	log.Printf("Request served. Elapsed: %dms", elapsed.Milliseconds())
	return elapsed.Milliseconds()
}
