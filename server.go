package main

import (
	"encoding/json"
	"fmt"
	. "jacksnake/models"
	"log"
	"net/http"
	"os"
	"time"
)

// Middleware

const ServerID = "battlesnake/github/starter-snake-go"

func withServerID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next(w, r)
	}
}

func withTimer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		fmt.Printf("%s", time.Since(t))
	}
}

func withMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return withTimer(withServerID(next))
}

// Start Battlesnake Server

type Responder interface {
	Info() BattlesnakeInfoResponse
	Start(gameState GameState)
	End(gameState GameState)
	Move(state GameState) BattlesnakeMoveResponse
}

func buildHandleIndex(responder Responder) func(http.ResponseWriter, *http.Request) {
	HandleIndex := func(w http.ResponseWriter, r *http.Request) {
		response := responder.Info()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("ERROR: Failed to encode info response, %s", err)
		}
	}
	return HandleIndex
}

func buildHandleStart(responder Responder) func(http.ResponseWriter, *http.Request) {
	handleStart := func(w http.ResponseWriter, r *http.Request) {
		state := GameState{}
		err := json.NewDecoder(r.Body).Decode(&state)
		if err != nil {
			log.Printf("ERROR: Failed to decode start json, %s", err)
			return
		}
		responder.Start(state)
	}
	return handleStart
	// Nothing to respond with here
}

func buildHandleMove(responder Responder) func(http.ResponseWriter, *http.Request) {
	handleMove := func(w http.ResponseWriter, r *http.Request) {
		state := GameState{}
		err := json.NewDecoder(r.Body).Decode(&state)
		if err != nil {
			log.Printf("ERROR: Failed to decode move json, %s", err)
			return
		}

		response := responder.Move(state)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("ERROR: Failed to encode move response, %s", err)
			return
		}
	}
	return handleMove
	// Nothing to respond with here
}

func buildHandleEnd(responder Responder) func(http.ResponseWriter, *http.Request) {
	handleEnd := func(w http.ResponseWriter, r *http.Request) {
		state := GameState{}
		err := json.NewDecoder(r.Body).Decode(&state)
		if err != nil {
			log.Printf("ERROR: Failed to decode end json, %s", err)
			return
		}

		responder.End(state)

		// Nothing to respond with here
	}
	return handleEnd
	// Nothing to respond with here
}

func RunServer(responder Responder) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	http.HandleFunc("/", withServerID(buildHandleIndex(responder)))
	http.HandleFunc("/start", withServerID(buildHandleStart(responder)))
	http.HandleFunc("/move", withServerID(buildHandleMove(responder)))
	http.HandleFunc("/end", withServerID(buildHandleEnd(responder)))

	log.Printf("Running Battlesnake at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
