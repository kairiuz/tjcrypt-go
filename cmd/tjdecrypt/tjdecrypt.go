package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.teknik.io/wobm/tjcrypt-go/pkg/tjcrypt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [input file]\n", os.Args[0])
		os.Exit(2)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	check(err)

	decrypted, err := tjcrypt.Decrypt(data)
	check(err)

	out := os.Stdout

	if len(os.Args) > 2 {
		out, err = os.Create(os.Args[2])
		check(err)
	}

	fmt.Fprint(out, string(decrypted))
}
