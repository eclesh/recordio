// Package recordio implements a file format for a sequence of
// records. It could be used to store, for example, serialized
// data structures to disk.
//
// Records are stored as an unsigned varint specifying the
// length of the data, and then the data itself as a binary blob.
//
// Example: reading
// 	f, _ := os.Open("file.dat")
// 	r := recordio.NewReader(f)
// 	for {
// 		data, err := r.Next()
// 		if err == io.EOF {
// 			break
// 		}
// 		// Do something with data
// 	}
// 	f.Close()
//
// Example: writing
// 	f, _ := os.Create("file.data")
// 	w := recordio.NewWriter(f)
// 	w.Write([]byte("this is a record"))
// 	w.Write([]byte("this is a second record"))
// 	f.Close()
//
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
		r.bufcap = uint64(size)
	}
	_, err = io.ReadFull(r.r, r.buf[:size])
	if err != nil {
		return nil, err
	}
	return r.buf[:size], nil
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
