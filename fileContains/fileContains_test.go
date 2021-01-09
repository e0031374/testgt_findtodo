package fileContains

import (
	"testing"
	"strings"
)

func TestNewStringTest(t *testing.T) {
	tests := []struct {
		input string
		expect bool
	}{
		{`"TODO"`, true},
		{`//"TODO"`, true},

		{`TODO`, false},
		{`"todo"`, false},
		{`"TO"`, false},
		{`"to be or not to be"`, false},
		{``, false},
	}

	for i, tt := range tests {
		todoTest := newStringTest(`"TODO"`)
		testResult := todoTest.Test(tt.input)
		if testResult != tt.expect {
			t.Fatalf("tests[%d] - failed match. input=%v expected=%v, got=%v",
                            i, tt.input, tt.expect, testResult)
		}
	}
}

