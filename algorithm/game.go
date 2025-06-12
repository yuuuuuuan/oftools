package algorithm

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// 生成一个4位不重复的数字字符串
func generateSecret() string {
	digits := []rune("0123456789")
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})
	return string(digits[:4])
}

// 比较 guess 和 secret，返回 A 和 B 的数量
func checkGuess(secret, guess string) (int, int) {
	A, B := 0, 0
	secretCount := make(map[rune]int)
	guessCount := make(map[rune]int)

	for i := 0; i < 4; i++ {
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

func isValidGuess(input string) bool {
	if len(input) != 4 {
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
	secret := generateSecret()
	attempts := 0

	fmt.Println("🔢 Welcome to the 4A Guessing Game!")
	fmt.Println("Try to guess the 4-digit number (no repeated digits).")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your guess: ")
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(input)

		if !isValidGuess(guess) {
			fmt.Println("⚠️  Invalid input! Please enter 4 unique digits.")
			continue
		}

		attempts++
		A, B := checkGuess(secret, guess)
		fmt.Printf("👉 %dA%dB\n", A, B)

		if A == 4 {
			fmt.Printf("🎉 Congratulations! The correct answer was %s. You solved it in %d attempts.\n", secret, attempts)
			break
		}
	}
	return nil
}
