package readerTest

import (
	"strings"
	"testing"
)

// tests for the ReaderTest.TestFn in the ReaderTest struct returned by HasSubstringTest
func TestHasSubstringTestFn(t *testing.T) {
	tests := []struct {
		input  string
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
		todoTest := HasSubstringTest(`"TODO"`).TestFn
		testResult := todoTest(tt.input)
		if testResult != tt.expect {
			t.Fatalf("tests[%d] - failed match. input=%v expected=%v, got=%v",
				i, tt.input, tt.expect, testResult)
		}
	}
}

func TestExactStringTestFn(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{`"TODO"`, true},

		{`//"TODO"`, false},
		{`TODO`, false},
		{`"todo"`, false},
		{`"TO"`, false},
		{`"to be or not to be"`, false},
		{``, false},
	}

	for i, tt := range tests {
		todoTest := ExactStringTest(`"TODO"`).TestFn
		testResult := todoTest(tt.input)
		if testResult != tt.expect {
			t.Fatalf("tests[%d] - failed match. input=%v expected=%v, got=%v",
				i, tt.input, tt.expect, testResult)
		}
	}
}

// tests the Test method of the ReaderTest struct
func TestReaderTestTestMethod(t *testing.T) {
	tests := []string{
		``,
		`TODO`,
		`Pen pineapple apple pen`,
	}

	trueStubTest := func(_ string) bool {
		return true
	}

	for i, tt := range tests {
		stubReaderTest := ReaderTest{DefaultVal: true, TestFn: trueStubTest}
		expected := true
		ttReader := strings.NewReader(tt)

		testResult := stubReaderTest.Test(ttReader)
		if testResult != expected {
			t.Fatalf("tests[%d] - failed match. input=%v expected=%v, got=%v",
				i, tt, expected, testResult)
		}
	}

	falseStubTest := func(_ string) bool {
		return false
	}
	for i, tt := range tests {
		stubReaderTest := ReaderTest{DefaultVal: false, TestFn: falseStubTest}
		expected := false
		ttReader := strings.NewReader(tt)

		testResult := stubReaderTest.Test(ttReader)
		if testResult != expected {
			t.Fatalf("tests[%d] - failed match. input=%v expected=%v, got=%v",
				i, tt, expected, testResult)
		}
	}
}
