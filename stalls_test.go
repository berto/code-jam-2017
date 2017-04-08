package main

import "testing"

var stallsTestOutput = `Case #1: 1 0
Case #2: 1 0
Case #3: 1 1
Case #4: 0 0
Case #5: 500 499
`

func Test_stalls(t *testing.T) {
	result := bathroomStalls(stallsTestInput)
	if result != stallsTestOutput {
		t.Error(result, " should equal", stallsTestOutput)
	}
}
