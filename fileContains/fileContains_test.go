package fileContains

import (
	"testing"
	"strings"
	"io/ioutil"
	"path/filepath"
	"os"
	"fmt"
)

type TestFile struct {
	contents string
	ans bool
}

func prepareTmpDir(tests []TestFile) (string, []string, error) {
	tmpDir, err := ioutil.TempDir("","")
	if err != nil {
		return tmpDir, nil, fmt.Errorf("Unable to create test dir%v", err)
	}

	filenames  := make([]string, len(tests))

	for i, tt := range tests {
		filename := "/test" + fmt.Sprint(i) + ".in"
		fp := filepath.Join(tmpDir, filename)
		f, err := os.Create(fp)
		if err != nil {
			return tmpDir, nil, fmt.Errorf("Unable to create test file %v", err)
		}
		_, err = f.WriteString(tt.contents)
		if err != nil {
			return tmpDir, nil, fmt.Errorf("Unable to write to test file %v", err)
		}

		f.Close()

		filenames[i] = filepath.Join(tmpDir, filename)
	}

	return tmpDir, filenames, nil
}

func TestNewStringTest(t *testing.T) {

	tests := []TestFile{
		{ `"TODO"`, true },
		{ `//"TODO"`, true  },
		{ `package main func main { "TODO" }`, true },
		{ `TODO`, false },
		{ `type Reader interface `, false },
	}
	target := `"TODO"`

	dir, filenames, err := prepareTmpDir(tests)
	defer os.RemoveAll(dir) // cleanup tmp dir
	if err != nil {
		t.Fatalf("setup file dir error")
	}

	for i, filename := range filenames {
		sb := new(strings.Builder)
		// https://stackoverflow.com/questions/13765797/the-best-way-to-get-a-string-from-a-writer
		// note need pointer for io.Writer, since it is the pointer of strings.Builder that implement the methods
		err := NewFileContainsWalkFunction(sb, target)(filename)
		if err != nil {
			t.Fatalf("tests[%d] - error testing file: %v ::\n", i, err)
		} else {
			filename = filename + "\n"
			writtenPath := sb.String()
			if tests[i].ans {
			    	if writtenPath != filename {
					t.Fatalf("tests[%d] - failed match.\nactualPath=%v,\nwrittenPath=%v,",
                        	    	i, filename, writtenPath)
				}
			} else {
			    	if writtenPath != "" {
					t.Fatalf("tests[%d] - failed match.\nactualPath=%v,\nwrittenPath=%v,",
                        	    	i, "[NIL]", writtenPath)
				}
			}
		}
	}
}
