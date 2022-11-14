package file_parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseData(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines")
		}
		line := s.Text()
		out = append(out, line)
	}
	return out, nil
}

// ParseDataFixed has the fixes for the problems
// found by fuzzing ParseData. Copy the specified lines to
// ParseData to see the fuzzing tests pass.
func ParseDataFixed(r io.Reader) ([]string, error) {
	// first line is number of lines
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}
	// test 1 fix start
	if count > 1000 {
		return nil, errors.New("too many")
	}
	// test 1 fix end
	// test 2 fix start
	if count < 0 {
		return nil, errors.New("no negative numbers")
	}
	// test 2 fix end
	out := make([]string, 0, count)
	// each following line is an entry in the slice
	for i := 0; i < count; i++ {
		hasLine := s.Scan()
		if !hasLine {
			return nil, errors.New("too few lines")
		}
		line := s.Text()
		// test 3 fix start
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			return nil, errors.New("blank line")
		}
		// test 3 fix end
		out = append(out, line)
	}
	// return the slice
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}
