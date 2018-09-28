package hangman

import (
	"database/sql"
	"log"
	// Use to access pgsql
	_ "github.com/lib/pq"
)

// CreateGame : Insert a new game into the database
func CreateGame(game Game) {
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
	db, err := sql.Open("potsgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Query(`INSERT INTO games (uuid, turns_left, word, used, available_hints) VALUES ($1, $2, $3, $4, $5)`,
		game.ID, game.TurnsLeft, game.Letters, game.Used, game.AvailableHints)
}
