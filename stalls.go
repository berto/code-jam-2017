package main

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	max        int
	min        int
	index      int
	difference int
}

func bathroomStalls(numbers string) (result string) {
	taskList := strings.Split(numbers, "\n")
	length, _ := strconv.Atoi(taskList[0])
	for i := 1; i < length+1; i++ {
		result += fmt.Sprintf("Case #%v: %s\n", i, findNextStall(strings.Split(taskList[i], " ")...))
	}
	return result
}

func findNextStall(stallsInfo ...string) string {
	stallsLength, _ := strconv.Atoi(stallsInfo[0])
	people, _ := strconv.Atoi(stallsInfo[1])
	stalls := make([]bool, stallsLength+2)
	stalls[0] = true
	stalls[stallsLength+1] = true
	p := position{0, 0, 0, 1}
	p = placePerson(stalls, people, p)
	return fmt.Sprintf("%v %v", p.max, p.min)
}

func placePerson(stalls []bool, people int, p position) position {
	if people == 0 {
		return p
	}
	p = position{0, 0, 0, 1}
	for i := 0; i < len(stalls); i++ {
		if !stalls[i] {
			left := findEmptyStalls(stalls[:i], "left")
			right := findEmptyStalls(stalls[i+1:], "right")
			difference := left - right
			if difference < 0 {
				difference = -difference
			}
			if difference == 1 || difference == 0 {
				max, min := findMaxAndMin(left, right)
				if max > p.max || (max == p.max && difference < p.difference) {
					p = position{max, min, i, difference}
				}
			}
		}
	}
	stalls[p.index] = true
	return placePerson(stalls, people-1, p)
}

func findEmptyStalls(stalls []bool, direction string) int {
	length := len(stalls)
	for i := 0; i < length; i++ {
		nextLeft := direction == "left" && stalls[length-i-1]
		nextRight := direction == "right" && stalls[i]
		if nextLeft || nextRight {
			return i
		}
	}
	return length
}

func findMaxAndMin(left, right int) (int, int) {
	if left > right {
		return left, right
	}
	return right, left
}
