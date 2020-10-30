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

This project is licensed under the MIT license. See LICENSE for more details.
