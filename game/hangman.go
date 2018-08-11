package hangman

import (
	"math/rand"
	"time"
)

// PickWord : Randomly get a word from a set of words.
func PickWord(words []string) string {
	rand.Seed(time.Now().Unix())
	wordIndex := rand.Intn(len(words))
	return words[wordIndex]
}

// LetterInWord : Check if word contains the given letter.
func LetterInWord(guess string, letters []string) bool {
	for _, letter := range letters {
		if guess == letter {
			return true
		}
	}
	return false
}

// RevealWord : reveal the word by checking if the guesses made
// are part of the choosen word.
func RevealWord(letters []string, used []string) string {
	revealedWord := ""

	for _, wordLetter := range letters {
		if LetterInWord(wordLetter, used) {
			revealedWord += wordLetter
		} else {
			revealedWord += "_"
		}
	}

	return revealedWord
}

// HasWon : Check if the player has won the game
func HasWon(letters []string, used []string) bool {
	ocurrences := 0
	for _, letter := range letters {
		for _, goodGuess := range used {
			if letter == goodGuess {
				ocurrences++
			}
		}
	}
	return ocurrences >= len(letters)
}