# recordio

Package recordio implements a file format for a sequence of
records. It could be used to store, for example, serialized
data structures to disk.

Records are stored as an unsigned varint specifying the
length of the data, and then the data itself as a binary blob.

## Example: reading with a Scanner
	f, _ := os.Open("file.data")
	defer f.Close()
	scanner := recordio.NewScanner(f)
	for scanner.Scan() {
		data := scanner.Bytes()
		// Do something with data
	}
	if err := scanner.Err(); err != nil {
		// Do error handling
	}

## Example: reading with a Reader
	f, _ := os.Open("file.dat")
	f.Close()
	r := recordio.NewReader(f)
	for {
		data, err := r.Next()
		if err == io.EOF {
			break
		}
		// Do something with data
	}

## Example: writing
	f, _ := os.Create("file.data")
	w := recordio.NewWriter(f)
	w.Write([]byte("this is a record"))
	w.Write([]byte("this is a second record"))
	f.Close()

## License

Copyright (C) 2012 Eric Lesh

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
