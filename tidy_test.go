package main

import "testing"

var tidyTestOutput = `Case #1: 129
Case #2: 999
Case #3: 7
Case #4: 99999999999999999
`

func Test_tidy(t *testing.T) {
	result := tidy(tidyTestInput)
	if result != tidyTestOutput {
		t.Error(result, " should equal", tidyTestOutput)
	}
}
