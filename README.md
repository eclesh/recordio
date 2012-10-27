# recordio

Package recordio implements a file format for a sequence of
records. It could be used to store, for example, serialized
data structures to disk.

Records are stored as an unsigned varint specifying the
length of the data, and then the data itself as a binary blob.

## Example: reading
	f, _ := os.Open("file.dat")
	r := recordio.NewReader(f)
	for {
		data, err := r.Next()
		if err == io.EOF {
			break
		}
		// Do something with data
	}
	f.Close()

## Example: writing
	f, _ := os.Create("file.data")
	w := recordio.NewWriter(f)
	w.Write([]byte("this is a record"))
	w.Write([]byte("this is a second record"))
	f.Close()
