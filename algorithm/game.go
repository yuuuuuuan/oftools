package algorithm

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generateSecret(length int) string {
	digits := []rune("0123456789")
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})
	return string(digits[:length])
}

func checkGuess(secret, guess string) (int, int) {
	A, B := 0, 0
	secretCount := make(map[rune]int)
	guessCount := make(map[rune]int)

	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			A++
		} else {
			secretCount[rune(secret[i])]++
			guessCount[rune(guess[i])]++
		}
	}

	for k, v := range guessCount {
		if sc, ok := secretCount[k]; ok {
			if v < sc {
				B += v
			} else {
				B += sc
			}
		}
	}
	return A, B
}

func isValidGuess(input string, length int) bool {
	if len(input) != length {
		return false
	}
	seen := make(map[rune]bool)
	for _, ch := range input {
		if ch < '0' || ch > '9' || seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

func GameWordle() error {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	// Select difficulty
	var difficulty string
	var length int

	for {
		fmt.Print("Choose difficulty (easy / normal / hard): ")
		difficultyInput, _ := reader.ReadString('\n')
		difficulty = strings.ToLower(strings.TrimSpace(difficultyInput))

		switch difficulty {
		case "easy":
			length = 4
		case "normal":
			length = 5
		case "hard":
			length = 6
		default:
			fmt.Println("⚠️ Invalid difficulty! Please type: easy, normal, or hard.")
			continue
		}
		break
	}

	secret := generateSecret(length)
	attempts := 0

	fmt.Printf("🔢 Game Start! Guess the %d-digit number (no repeated digits).\n", length)

	for {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(input)

		if !isValidGuess(guess, length) {
			fmt.Printf("⚠️ Invalid input! Please enter %d unique digits.\n", length)
			continue
		}

		attempts++
		A, B := checkGuess(secret, guess)
		fmt.Printf("👉 %dA%dB\n", A, B)

		if A == length {
			fmt.Printf("🎉 Congratulations! The correct answer was %s. You solved it in %d attempts.\n", secret, attempts)
			break
		}
	}
	return nil
}
