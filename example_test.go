package recordio_test

import (
	"fmt"
	"log"
	"os"

	"github.com/eclesh/recordio"
)

func ExampleScanner() {
	f, err := os.Open("file.dat")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := recordio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func ExampleWriter() {
	f, err := os.Create("file.dat")
	if err != nil {
		log.Fatalln(err)
	}
	w := recordio.NewWriter(f)
	w.Write([]byte("this is a record"))
	w.Write([]byte("this is a second record"))
	f.Close()
}
