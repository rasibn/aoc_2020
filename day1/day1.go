package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Main() {
	numbers, err := readNumberList("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, number3 := range numbers {
		target := 2020 - number3
		hashmap := make(map[int]int)
		for _, number2 := range numbers {
			leftover := target - number2

			if number2 == number3 || target == number3 {
				continue
			}

			if number2 == 0 || target == 0 || number3 == 0 {
				continue
			}

			if leftover < 0 {
				continue
			}
			if _, found := hashmap[number2]; found {
				fmt.Println(number2 * number3 * leftover)
				os.Exit(0)
			}
			hashmap[leftover] = number2
		}
	}
}

func readNumberList(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	array := make([]int, 0)

	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		array = append(array, number)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return array, nil
}
