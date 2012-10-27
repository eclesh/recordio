package recordio

import (
	"bytes"
	"io"
	"testing"
)

func TestEmpty(t *testing.T) {
	buf := new(bytes.Buffer)
	r := NewReader(buf)
	if _, err := r.Next(); err != io.EOF {
		t.Fatalf("got %v, expected %v", err, io.EOF)
	}
}

func TestOrder(t *testing.T) {
	testValues := []string{
		"first",
		"second",
		"third",
	}

	buf := new(bytes.Buffer)

	w := NewWriter(buf)
	for _, val := range testValues {
		_, err := w.Write([]byte(val))
		if err != nil {
			t.Fatalf("unexpected error %v for value '%s'", err, val)
		}
	}

	r := NewReader(buf)
	for _, expected := range testValues {
		received, err := r.Next()
		if err != nil {
			t.Fatalf("unexpected error '%v' for value '%s'", err, expected)
		}
		if len(received) != len(expected) {
			t.Fatalf("expected length %d, got %d for value '%s'", len(expected), len(received), expected)
		}
		if string(received) != expected {
			t.Fatalf("expected '%s', got '%s'", expected, received)
		}
	}
}
