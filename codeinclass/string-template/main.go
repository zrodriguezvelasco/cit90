package main

import (
	"strings"
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	// CREATE STRING
	// assign to variable
	name := "Zoraida"
	str := `html here ` + name + ` html here`
	fmt.Println(str)

	// CREATE STRING
	// string print
	s := fmt.Sprint(`blah ` + name + ` blah`)
	fmt.Print(s)

	// CREATE FILE
	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}
	
	// io.Copy to File
	io.Copy(nf, strings.NewReader(s))

	// CREATE FILE
	// writestring to file
	nf2, err := os.Create("newfile2.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("whoops2", err)
	}

	fmt.Println("bytes", n)
}