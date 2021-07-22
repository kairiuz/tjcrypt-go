package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git.teknik.io/wobm/tjcrypt-go/pkg/tjcrypt"
)

const TJ_DEFAULT_KEY = "\xd2\"\x82\x7f\xe9\xd3r\xa5$\x90\x8cm\n\x96\xcb\xa3"
const TJ_PASSWORD = "MODDEDBYPL0NK!!!"

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
