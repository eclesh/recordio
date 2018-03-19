// Package recordio implements a file format for a sequence of
// records. It could be used to store, for example, serialized
// data structures to disk.
//
// Records are stored as an unsigned varint specifying the
// length of the data, and then the data itself as a binary blob.
package recordio

import (
	"bufio"
	"encoding/binary"
	"io"
)

type Reader struct {
	r      io.Reader // the reader
	buf    []byte    // the buffer
	bufcap uint64    // the capacity of the buffer
}

// NewReader returns a new reader. If r doesn't implement
// io.ByteReader, it will be wrapped in a bufio.Reader.
func NewReader(r io.Reader) *Reader {
	if _, ok := r.(io.ByteReader); !ok {
		r = bufio.NewReader(r)
	}
	return &Reader{r: r}
}

// Next returns the next data record. It returns io.EOF if there are
// no more records.
func (r *Reader) Next() ([]byte, error) {
	var err error
	size, err := binary.ReadUvarint(r.r.(io.ByteReader))
	if err != nil {
		return nil, err
	}
	if size > r.bufcap {
		r.buf = make([]byte, size)
		r.bufcap = size
	}
	_, err = io.ReadFull(r.r, r.buf[:size])
	if err != nil {
		return nil, err
	}
	return r.buf[:size], nil
}

// A Scanner is a convenient method for reading records sequentially.
type Scanner struct {
	r       io.Reader // the reader
	err     error
	buf     []byte
	bufsize uint64
	bufcap  uint64
}

// NewScanner creates a new Scanner from reader r.
func NewScanner(r io.Reader) *Scanner {
	if _, ok := r.(io.ByteReader); !ok {
		r = bufio.NewReader(r)
	}
	return &Scanner{r: r}
}

// Scan chugs through the input record by record and stops at the first
// error or EOF.
func (s *Scanner) Scan() bool {
	size, err := binary.ReadUvarint(s.r.(io.ByteReader))
	if err != nil {
		s.err = err
		return false
	}
	s.bufsize = size
	if size > s.bufcap {
		s.buf = make([]byte, size)
		s.bufcap = size
	}
	_, err = io.ReadFull(s.r, s.buf[:size])
	if err != nil {
		s.err = err
		return false
	}
	return true
}

// Bytes returns the most recently scanned record.
func (s *Scanner) Bytes() []byte {
	return s.buf[:s.bufsize]
}

// Err returns the most recent error or nil if the error was EOF.
func (s *Scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

type Writer struct {
	w io.Writer // the writer
}

// NewWriter returns a new writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
	}
}

// Write writes a data record.
func (w *Writer) Write(data []byte) (int, error) {
	size := uint64(len(data))
	max_size := size + binary.MaxVarintLen64
	buf := make([]byte, max_size)
	n := binary.PutUvarint(buf, size)
	buf = append(buf[:n], data...)
	return w.w.Write(buf)
}
