package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Проверяем, что ввод является математической операцией
	if !isValidExpression(input) {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	// Проверяем, что используется только одна система счисления
	if isMixedNumeralSystem(input) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	// Парсим числа и оператор
	num1, num2, operator := parseExpression(input)

	// Выполняем операцию
	result := calculate(num1, num2, operator)

	// Выводим результат
	// if isRomanNumeral(num1) {
	// 	fmt.Println(romanToArabic(result))
	// } else {
	// 	fmt.Println(result)
	// }
	if strings.Contains(input, "I") || strings.Contains(input, "V") || strings.Contains(input, "X") {
		fmt.Println(romanToArabic(strconv.Itoa(result)))
	} else {
		fmt.Println(result)
	}
}

// Проверяем, что ввод является математической операцией
func isValidExpression(expression string) bool {
	parts := strings.Split(expression, " ")
	return len(parts) == 3 && (parts[1] == "+" || parts[1] == "-" || parts[1] == "*" || parts[1] == "/")
}

// Проверяем, что используется только одна система счисления
func isMixedNumeralSystem(expression string) bool {
	parts := strings.Split(expression, " ")
	return (strings.Contains(parts[0], "I") || strings.Contains(parts[0], "V") || strings.Contains(parts[0], "X")) &&
		(!strings.Contains(parts[2], "I") && !strings.Contains(parts[2], "V") && !strings.Contains(parts[2], "X")) ||
		(!strings.Contains(parts[0], "I") && !strings.Contains(parts[0], "V") && !strings.Contains(parts[0], "X")) &&
			(strings.Contains(parts[2], "I") || strings.Contains(parts[2], "V") || strings.Contains(parts[2], "X"))
}

// Парсим числа и оператор
func parseExpression(expression string) (int, int, string) {
	parts := strings.Split(expression, " ")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[2])
	operator := parts[1]

	// Преобразуем римские числа в арабские
	if isRomanNumeral(parts[0]) {
		num1 = romanToArabic(parts[0])
	}
	if isRomanNumeral(parts[2]) {
		num2 = romanToArabic(parts[2])
	}

	return num1, num2, operator
}

// Выполняем операцию
func calculate(num1 int, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Выдача паники, так как деление на ноль невозможно")
		}
		return num1 / num2
	default:
		panic("Неизвестный оператор")
	}
}

// Проверяем, что число является римским
func isRomanNumeral(number string) bool {
	return number == "I" || number == "II" || number == "III" || number == "IV" || number == "V" || number == "VI" || number == "VII" || number == "VIII" || number == "IX" || number == "X"
}

// Преобразуем римское число в арабское
func romanToArabic(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		panic("Неверное римское число")
	}
}
