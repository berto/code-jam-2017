package main

import "testing"

var pancakeTestOutput = `Case #1: IMPOSSIBLE
Case #2: 3
Case #3: 0
Case #4: IMPOSSIBLE
`

func Test_pancake(t *testing.T) {
	result := pancakeFlipper(pancakeTestInput)
	if result != pancakeTestOutput {
		t.Error(result, " should equal", pancakeTestOutput)
	}
}
