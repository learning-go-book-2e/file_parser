package file_parser

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParseData(t *testing.T) {
	data := []struct {
		name   string
		in     []byte
		out    []string
		errMsg string
	}{
		{
			name:   "simple",
			in:     []byte("3\nhello\ngoodbye\ngreetings\n"),
			out:    []string{"hello", "goodbye", "greetings"},
			errMsg: "",
		},
		{
			name:   "empty_error",
			in:     []byte(""),
			out:    nil,
			errMsg: "empty",
		},
		{
			name:   "zero",
			in:     []byte("0\n"),
			out:    []string{},
			errMsg: "",
		},
		{
			name:   "number_error",
			in:     []byte("asdf\nhello\ngoodbye\ngreetings\n"),
			out:    nil,
			errMsg: `strconv.Atoi: parsing "asdf": invalid syntax`,
		},
		{
			name:   "line_count_error",
			in:     []byte("4\nhello\ngoodbye\ngreetings\n"),
			out:    nil,
			errMsg: "too few lines",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			r := bytes.NewReader(d.in)
			out, err := ParseData(r)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if diff := cmp.Diff(d.out, out); diff != "" {
				t.Error(diff)
			}
			if diff := cmp.Diff(d.errMsg, errMsg); diff != "" {
				t.Error(diff)
			}
			//if err == nil {
			//	roundTrip := ToData(out)
			//	if diff := cmp.Diff(d.in, roundTrip); diff != "" {
			//		t.Error(diff)
			//	}
			//}
		})
	}
}

func FuzzParseData(f *testing.F) {
	testcases := [][]byte{
		[]byte("3\nhello\ngoodbye\ngreetings\n"),
		[]byte("0\n"),
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		r := bytes.NewReader(in)
		out, err := ParseData(r)
		if err != nil {
			t.Skip("invalid number")
		}
		roundTrip := ToData(out)
		rtr := bytes.NewReader(roundTrip)
		out2, err := ParseData(rtr)
		if diff := cmp.Diff(out, out2); diff != "" {
			t.Error(diff)
		}
	})
}
