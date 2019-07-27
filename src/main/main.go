package main

import (
	"log"
	"os"
	"test/ch"
)

func init() {
	log.SetOutput(os.Stdout)
}

// https://javascript.ctolib.com/article/compares/93347
func main() {
	ch.Main()
}
