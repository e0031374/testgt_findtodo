package readerTest

import (
	"bufio"
	"io"
	"strings"
)

// log use of bufio.NewScanner to read from file
// log this bit about using bufio.ScanWords
// https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4

// default will test against empty string

type ReaderTest struct {
	DefaultVal bool
	TestFn     func(string) bool
}

func (st ReaderTest) Test(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		input := sc.Text()
		if st.TestFn(input) {
			return st.TestFn(input)
		}
	}
	return st.DefaultVal
}

func HasSubstringTest(substr string) ReaderTest {
	testFn := func(input string) bool {
		return strings.Contains(input, substr)
	}
	return ReaderTest{
		DefaultVal: false,
		TestFn:     testFn,
	}
}

func ExactStringTest(substr string) ReaderTest {
	testFn := func(input string) bool {
		return input == substr
	}
	return ReaderTest{
		DefaultVal: false,
		TestFn:     testFn,
	}
}
