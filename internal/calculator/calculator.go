package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(expression string) (float64, error) {
	result, err := Calc(expression)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate: %v", err)
	}
	return result, nil
}

// Функция для преобразования строки в float64
func StringToFloat64(str string) float64 {
	degree := float64(1)
	var res float64 = 0
	var invers bool = false
	for i := len(str); i > 0; i-- {
		if str[i-1] == '-' {
			invers = true
		} else {
			res += float64(9-int('9'-str[i-1])) * degree
			degree *= 10
		}
	}
	if invers {
		res = 0 - res
	}
	return res
}

// Проверка, является ли символ знаком операции
func IsSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

// Основная функция для вычисления выражений
func Calc(expression string) (float64, error) {
	// Убираем пробелы
	expression = strings.ReplaceAll(expression, " ", "")
	if len(expression) < 3 {
		return 0, fmt.Errorf("invalid expression")
	}
	for {
		openIdx := strings.LastIndex(expression, "(")
		if openIdx == -1 {
			break
		}
		closeIdx := strings.Index(expression[openIdx:], ")")
		if closeIdx == -1 {
			return 0, fmt.Errorf("mismatched parentheses")
		}
		innerExpr := expression[openIdx+1 : openIdx+closeIdx]
		result, err := Calc(innerExpr)
		if err != nil {
			return 0, err
		}
		// Заменяем выражение в скобках на результат
		expression = expression[:openIdx] + fmt.Sprintf("%f", result) + expression[openIdx+closeIdx+1:]
	}

	operators := []rune{'*', '/'}
	for _, op := range operators {
		for {
			opIdx := strings.IndexRune(expression, op)
			if opIdx == -1 {
				break
			}
			leftIdx := opIdx - 1
			for leftIdx >= 0 && !IsSign(rune(expression[leftIdx])) {
				leftIdx--
			}
			leftIdx++

			rightIdx := opIdx + 1
			for rightIdx < len(expression) && !IsSign(rune(expression[rightIdx])) {
				rightIdx++
			}

			leftOperand := expression[leftIdx:opIdx]
			rightOperand := expression[opIdx+1 : rightIdx]
			leftNum, err := strconv.ParseFloat(leftOperand, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", leftOperand)
			}
			rightNum, err := strconv.ParseFloat(rightOperand, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", rightOperand)
			}
			var result float64
			switch op {
			case '*':
				result = leftNum * rightNum
			case '/':
				if rightNum == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = leftNum / rightNum
			}
			expression = expression[:leftIdx] + fmt.Sprintf("%f", result) + expression[rightIdx:]
		}
	}

	for _, op := range []rune{'+', '-'} {
		for {
			opIdx := strings.IndexRune(expression, op)
			if opIdx == -1 {
				break
			}
			leftIdx := opIdx - 1
			for leftIdx >= 0 && !IsSign(rune(expression[leftIdx])) {
				leftIdx--
			}
			leftIdx++

			rightIdx := opIdx + 1
			for rightIdx < len(expression) && !IsSign(rune(expression[rightIdx])) {
				rightIdx++
			}
			leftOperand := expression[leftIdx:opIdx]
			rightOperand := expression[opIdx+1 : rightIdx]
			leftNum, err := strconv.ParseFloat(leftOperand, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", leftOperand)
			}
			rightNum, err := strconv.ParseFloat(rightOperand, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", rightOperand)
			}

			// Выполняем операцию
			var result float64
			switch op {
			case '+':
				result = leftNum + rightNum
			case '-':
				result = leftNum - rightNum
			}
			expression = expression[:leftIdx] + fmt.Sprintf("%f", result) + expression[rightIdx:]
		}
	}

	// Преобразуем финальный результат в число
	finalResult, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid expression result")
	}
	return finalResult, nil
}
