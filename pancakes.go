package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pancakeFlipper(tasks string) (result string) {
	taskList := strings.Split(tasks, "\n")
	length, _ := strconv.Atoi(taskList[0])
	for i := 1; i < length+1; i++ {
		result += fmt.Sprintf("Case #%v: %s\n", i, flip(strings.Split(taskList[i], " ")...))
	}
	return result
}

func flip(pancakes ...string) string {
	size, _ := strconv.Atoi(pancakes[1])
	record := make(map[string]bool)
	result := oversizeFlip(pancakes[0], size, 0, record)
	return result
}

func oversizeFlip(pancakes string, size int, flips int, record map[string]bool) string {
	if record[pancakes] {
		return "IMPOSSIBLE"
	}
	if isHappy(pancakes) {
		return strconv.Itoa(flips)
	}
	record[pancakes] = true
	flipping := false
	j := 0
	length := len(pancakes)
	for i := 0; i < length; i++ {
		if !flipping && length-(i+1) < size-1 {
			break
		}
		if string(pancakes[i]) == "-" {
			flipping = true
		}
		if flipping && j < size {
			side := '+'
			if string(pancakes[i]) == "+" {
				side = '-'
			}
			pancakes = flipPancake(pancakes, side, i)
			j++
		}
	}
	return oversizeFlip(pancakes, size, flips+1, record)
}

func isHappy(pancakes string) bool {
	for i := 0; i < len(pancakes); i++ {
		if string(pancakes[i]) != "+" {
			return false
		}
	}
	return true
}

func flipPancake(pancakes string, s rune, i int) string {
	list := []rune(pancakes)
	list[i] = s
	return string(list)
}
