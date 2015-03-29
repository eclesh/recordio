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

func TestScanner(t *testing.T) {
	testValues := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}

	buf := new(bytes.Buffer)

	w := NewWriter(buf)
	for _, val := range testValues {
		_, err := w.Write(val)
		if err != nil {
			t.Fatalf("unexpected error %v for value '%s'", err, val)
		}
	}

	scanner := NewScanner(buf)
	i := 0
	for scanner.Scan() {
		if i >= len(testValues) {
			t.Fatalf("scanner scanned for %d elements but only %d exist", i, len(testValues))
		}
		expected := testValues[i]
		data := scanner.Bytes()
		if !bytes.Equal(data, expected) {
			t.Fatalf("expected value '%s', got '%s' instead", string(expected), string(data))
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("unexpected error %v during scan", err)
	}
}
