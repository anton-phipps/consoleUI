package main

import (
	"log"
	"os"

	"golang.org/x/term"
)

func goGetChar() byte {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var b [1]byte
	os.Stdin.Read(b[:])
	return b[0]
}
