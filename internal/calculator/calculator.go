package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

// Calculate — внешний интерфейс калькулятора
func Calculate(expression string) (float64, error) {
	result, err := Calc(expression)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate: %v", err)
	}
	return result, nil
}

// IsSign — проверка, является ли символ знаком операции
func IsSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

// Calc — рекурсивный парсер и вычислитель выражений с поддержкой скобок
func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if len(expression) < 3 {
		return 0, fmt.Errorf("invalid expression")
	}

	// Обработка скобок
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
		expression = expression[:openIdx] + fmt.Sprintf("%f", result) + expression[openIdx+closeIdx+1:]
	}

	// *, /
	for _, op := range []rune{'*', '/'} {
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
			if op == '*' {
				result = leftNum * rightNum
			} else {
				if rightNum == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = leftNum / rightNum
			}
			expression = expression[:leftIdx] + fmt.Sprintf("%f", result) + expression[rightIdx:]
		}
	}

	// +, -
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
			var result float64
			if op == '+' {
				result = leftNum + rightNum
			} else {
				result = leftNum - rightNum
			}
			expression = expression[:leftIdx] + fmt.Sprintf("%f", result) + expression[rightIdx:]
		}
	}

	finalResult, err := strconv.ParseFloat(expression, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid expression result")
	}
	return finalResult, nil
}
