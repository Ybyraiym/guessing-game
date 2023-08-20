package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Определение диапазона минимального и максимального чисел
	min, max := 1, 100

	// Установка семени для генератора случайных чисел на основе текущего времени
	rand.Seed(time.Now().UnixNano())

	// Генерация случайного числа в заданном диапазоне
	secretNumber := rand.Intn(max-min) + min

	// Вывод инструкций для игрока
	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Please input your guess")

	// Счетчик попыток
	attempts := 0

	// Бесконечный цикл для игры
	for {
		attempts++

		// Создание объекта для чтения ввода из терминала
		reader := bufio.NewReader(os.Stdin)

		// Чтение введенной строки до символа новой строки
		input, err := reader.ReadString('\n')

		// Обработка ошибок чтения ввода
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}

		// Удаление символа новой строки из введенной строки
		input = strings.TrimSpace(input)

		// Конвертация введенной строки в целое число
		guess, err := strconv.Atoi(input)

		// Обработка ошибок конвертации
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}

		// Вывод введенной игроком догадки
		fmt.Println("Your guess is", guess)

		// Сравнение догадки игрока с секретным числом
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Try again")
		} else {
			// Вывод поздравления при угадывании числа и количества попыток
			fmt.Println("Correct, you Legend! You guessed right after", attempts, "attempts")
			break
		}
	}
}
