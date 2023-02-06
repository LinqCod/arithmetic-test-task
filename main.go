package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var nums = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
var neededResult = 200

// main Основная идея в применении комбинаторики и систем счисления:
// три возможных оператора (+ / - / nothing) -> работаем с числами в троичной системе
// всего 3^9 вариаций позиционирования операторов, пройдемся по ним, используя 9-значные числа в троичной системе в качестве последовательностей операторов
// пример: [9,8,7,6,5,4,3,2,1,0] с последовательностью операторов: 212100122, эквивалентно: 98-76-5+4+3-210
// Соответственно ('0' == '+', '1' == '-', '2' == ”)
func main() {
	// крайние границы последовательностей операторов
	currentOperatorsSequence := convertNumberFrom3To10System(100000000)
	maxOperatorsSequence := convertNumberFrom3To10System(222222222)

	for currentOperatorsSequence <= maxOperatorsSequence {
		if expression, correct := getResultArithmeticExpression(int64(currentOperatorsSequence)); correct {
			fmt.Println(expression)
		}
		currentOperatorsSequence++
	}

}

// getResultArithmeticExpression функция получения строкового представления арифметического выражения и флага,
// говорящего об удовлетворении результата условию
func getResultArithmeticExpression(seq int64) (string, bool) {
	operators := []rune(strconv.FormatInt(seq, 3))
	res := 0

	currentIndex := 0
	currentSummeryNumber := nums[currentIndex]
	lastArithmeticOperator := '0'

	for currentIndex < len(operators) {
		if operators[currentIndex] == '2' {
			currentSummeryNumber = currentSummeryNumber*10 + nums[currentIndex+1]
		} else {
			res += calcNumberWithArithmeticOperator(lastArithmeticOperator, currentSummeryNumber)
			lastArithmeticOperator = operators[currentIndex]

			currentSummeryNumber = nums[currentIndex+1]
		}

		currentIndex++
	}
	res += calcNumberWithArithmeticOperator(lastArithmeticOperator, currentSummeryNumber)

	if res == neededResult {
		return createResultString(operators), true
	}

	return "", false
}

// createResultString функция создания результирующего арифметического выражения
func createResultString(operators []rune) string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(nums[0]))

	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '0':
			builder.WriteRune('+')
		case '1':
			builder.WriteRune('-')
		}
		builder.WriteString(strconv.Itoa(nums[i+1]))
	}

	return builder.String()
}

// calcNumberWithArithmeticOperator функция для определения знака числа по стоящему перед ним оператору
func calcNumberWithArithmeticOperator(operator rune, number int) int {
	if operator == '0' {
		return number
	}

	return -number
}

// convertNumberFrom3To10System функция конвертации числа из троичной системы счисления в десятичную
func convertNumberFrom3To10System(num int) int {
	res := 0
	index := 0

	for num > 0 {
		lastDigit := num % 10
		res += lastDigit * int(math.Pow(3, float64(index)))
		index++
		num /= 10
	}

	return res
}
