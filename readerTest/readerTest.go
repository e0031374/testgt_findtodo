package readerTest

import (
	"bufio"
	"io"
	"strings"
)

// prevents ambiguous case if there is nothing to test
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
