package main

import (
	"log"
	"os"
	"test/ref"
)

func init() {
	log.SetOutput(os.Stdout)
}

// https://javascript.ctolib.com/article/compares/93347
func main() {
	ref.Main()
}
