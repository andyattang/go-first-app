package main

import (
	"log"
	"os"
	"web"
)

func init() {
	log.SetOutput(os.Stdout)
}

// https://javascript.ctolib.com/article/compares/93347
func main() {
	web.Main()
}
