package main

import (
	"encoding/json"
	"go-hangman/game"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type gameInfoJSON struct {
	ID             string   `json:"id"`
	TurnsLeft      int      `json:"turns_left"`
	Used           []string `json:"used"`
	AvailableHints int      `json:"available_hints"`
}

func newGame(w http.ResponseWriter, r *http.Request) {
	words := []string{
		"apple",
		"banana",
		"orange",
	}
	choosenWord := hangman.PickWord(words)
	game := hangman.NewGame(3, choosenWord)
	w.Header().Set("Location", strings.Join([]string{r.Host, "games", game.ID}, "/"))
}

func getGameInfo(w http.ResponseWriter, r *http.Request) {
	words := []string{
		"apple",
		"banana",
		"orange",
	}
	choosenWord := hangman.PickWord(words)
	game := hangman.NewGame(3, choosenWord)
	responseJSON := gameInfoJSON{
		ID:             game.ID,
		TurnsLeft:      game.TurnsLeft,
		Used:           game.Used,
		AvailableHints: game.AvailableHints,
	}
	buff, error := json.MarshalIndent(responseJSON, "", "\t")
	if error != nil {
		log.Fatal("Could not serialize game")
	}
	w.Write(buff)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", newGame).Methods("GET")
	router.HandleFunc("/games/{id}", getGameInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
