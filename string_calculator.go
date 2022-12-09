package tdd_golang

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}

	delimiter := ","
	inputs := strings.Fields(input)

	firstLine := inputs[0]
	if len(firstLine) >= 3 && firstLine[:2] == "//" {
		delimiter = firstLine[2:]
		inputs = inputs[1:]
	}

	numbers := []string{}
	for _, v := range inputs {
		numbers = append(numbers, strings.Split(v, delimiter)...)
	}

	sum := 0
	for _, v := range numbers {
		if v == "" {
			continue
		}
		number, err := strconv.Atoi(v)
		if err != nil {
			return 0, nil
		}
		if number < 0 {
			return 0, fmt.Errorf("negatives not allowed: %d", number)
		}
		if number > 1000 {
			continue
		}
		sum += number
	}
	return sum, nil
}
