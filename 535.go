package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
	16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
	21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
	26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
	31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
	36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
	41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
	46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
	51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV",
	56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
	61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
	66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
	71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV",
	76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
	81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
	86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
	91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
	96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
}

func isRoman(s string) bool {
	_, ok := romanToArabic[s]
	return ok
}

func isArabic(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return num >= 1 && num <= 10
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("неверное действие")
	}
}

func arabicToRomanResult(result int) (string, error) {
	if result <= 0 {
		return "", errors.New("результат меньше или равен нулю")
	}
	return arabicToRoman[result], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите запрос: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Ошибка ввода")
	}

	a, operator, b := parts[0], parts[1], parts[2]

	var isRomanInput bool
	if isRoman(a) && isRoman(b) {
		isRomanInput = true
	} else if isArabic(a) && isArabic(b) {
		isRomanInput = false
	} else {
		panic("Смешанные или неверные типы чисел")
	}

	var num1, num2 int
	if isRomanInput {
		num1 = romanToArabic[a]
		num2 = romanToArabic[b]
	} else {
		var err error
		num1, err = strconv.Atoi(a)
		if err != nil {
			panic("Неверное число")
		}
		num2, err = strconv.Atoi(b)
		if err != nil {
			panic("Неверное число")
		}
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		panic(err)
	}

	if isRomanInput {
		romanResult, err := arabicToRomanResult(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
