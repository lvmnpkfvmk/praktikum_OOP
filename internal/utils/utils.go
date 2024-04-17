package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Статический класс для вывода текста с задержкой

// Метод для выбора числа
func ChoiceNumber(max int, needCancellation bool) int {
	if max <= 0 {
		panic("Максимальное число должно быть больше 0")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "0" && needCancellation {
			return 0
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Введено не число", 10, 20)
			continue
		}

		if num < 1 || num > max {
			fmt.Print(fmt.Sprintf("Введено число, которое не входит в диапазон от 1 до %d. Повторите ввод:", max), 10, 20)
			continue
		}

		return num
	}
}