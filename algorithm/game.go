package algorithm

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
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

type Board [4][4]int

func Game2048() error {
	rand.Seed(time.Now().UnixNano())
	board := Board{}
	addRandom(&board)
	addRandom(&board)

	// 初始化键盘输入
	if err := keyboard.Open(); err != nil {
		return fmt.Errorf("无法打开键盘输入: %v", err)
	}
	defer keyboard.Close()

	for {
		clearScreen()
		printBoard(board)

		// 等待方向键输入
		fmt.Println("请输入方向 (↑ ↓ ← →) 或 Q 退出:")
		char, key, err := keyboard.GetKey()
		if err != nil {
			return fmt.Errorf("键盘输入失败: %v", err)
		}

		// 用户退出
		if char == 'q' || char == 'Q' {
			fmt.Println("感谢游戏，再见！")
			return nil
		}

		var moved bool
		switch key {
		case keyboard.KeyArrowUp:
			moved = moveUp(&board)
		case keyboard.KeyArrowDown:
			moved = moveDown(&board)
		case keyboard.KeyArrowLeft:
			moved = moveLeft(&board)
		case keyboard.KeyArrowRight:
			moved = moveRight(&board)
		default:
			// 如果不是方向键，忽略
			continue
		}

		if moved {
			addRandom(&board)
		}

		if checkGameOver(board) {
			clearScreen()
			printBoard(board)
			fmt.Println("游戏结束！")
			return nil
		}
		if checkWin(board) {
			clearScreen()
			printBoard(board)
			fmt.Println("恭喜你赢了！")
			return nil
		}
	}
}

func clearScreen() {
	// 通过打印控制符清屏
	fmt.Print("\033[H\033[2J")
}

func printBoard(board Board) {
	fmt.Println("  2048 游戏")
	fmt.Println("当前得分: ", getScore(board))

	// 打印棋盘
	for _, row := range board {
		fmt.Print(" ")
		for _, val := range row {
			if val == 0 {
				fmt.Print(" .  ")
			} else {
				fmt.Printf("%-4d", val)
			}
		}
		fmt.Println()
	}

	// 提示信息
	fmt.Println("\n[↑] 上 [↓] 下 [←] 左 [→] 右 [Q] 退出")
}

func getScore(board Board) int {
	score := 0
	for _, row := range board {
		for _, val := range row {
			score += val
		}
	}
	return score
}

func addRandom(board *Board) {
	var emptyCells []struct{ x, y int }
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if board[y][x] == 0 {
				emptyCells = append(emptyCells, struct{ x, y int }{x, y})
			}
		}
	}

	if len(emptyCells) == 0 {
		return
	}

	cell := emptyCells[rand.Intn(len(emptyCells))]
	board[cell.y][cell.x] = 2 * (rand.Intn(2) + 1)
}

func moveUp(board *Board) bool {
	return slide(board, func(i, j int) (int, int) { return i - 1, j })
}

func moveDown(board *Board) bool {
	return slide(board, func(i, j int) (int, int) { return i + 1, j })
}

func moveLeft(board *Board) bool {
	return slide(board, func(i, j int) (int, int) { return i, j - 1 })
}

func moveRight(board *Board) bool {
	return slide(board, func(i, j int) (int, int) { return i, j + 1 })
}

func slide(board *Board, getNextPos func(i, j int) (int, int)) bool {
	moved := false
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if board[i][j] == 0 {
				continue
			}

			x, y := j, i
			for {
				nextX, nextY := getNextPos(x, y)
				if nextX < 0 || nextX >= 4 || nextY < 0 || nextY >= 4 {
					break
				}

				if board[nextY][nextX] == 0 {
					board[nextY][nextX] = board[y][x]
					board[y][x] = 0
					x, y = nextX, nextY
					moved = true
				} else if board[nextY][nextX] == board[y][x] {
					board[nextY][nextX] *= 2
					board[y][x] = 0
					moved = true
					break
				} else {
					break
				}
			}
		}
	}
	return moved
}

func checkGameOver(board Board) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if board[y][x] == 0 {
				return false
			}
			if x+1 < 4 && board[y][x] == board[y][x+1] {
				return false
			}
			if y+1 < 4 && board[y][x] == board[y+1][x] {
				return false
			}
		}
	}
	return true
}

func checkWin(board Board) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if board[y][x] == 2048 {
				return true
			}
		}
	}
	return false
}
