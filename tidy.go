package main

import (
	"fmt"
	"strconv"
	"strings"
)

func tidy(numbers string) (result string) {
	taskList := strings.Split(numbers, "\n")
	length, _ := strconv.Atoi(taskList[0])
	for i := 1; i < length+1; i++ {
		result += fmt.Sprintf("Case #%v: %s\n", i, findTidyNumber(taskList[i]))
	}
	return result
}

func findTidyNumber(number string) string {
	ok, subtraction := isTidy(number)
	if ok {
		return number
	}
	intNumber, _ := strconv.Atoi(number)
	return findTidyNumber(strconv.Itoa(intNumber - subtraction))
}

func isTidy(number string) (bool, int) {
	length := len(number)
	for i := 0; i < length; i++ {
		if i+1 != length {
			current, _ := strconv.Atoi(string(number[i]))
			next, _ := strconv.Atoi(string(number[i+1]))
			if current > next {
				numberList := strings.Split(number, "")
				subtraction, _ := strconv.Atoi(strings.Join(numberList[i+1:], ""))
				return false, subtraction + 1
			}
		}
	}
	return true, 0
}
