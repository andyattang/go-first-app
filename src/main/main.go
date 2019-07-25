package main

import (
	"log"
	"os"
	"test/inter"
)

func init() {
	log.SetOutput(os.Stdout)
}

// https://javascript.ctolib.com/article/compares/93347
func main() {
	inter.Main()
}
