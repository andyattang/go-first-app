package search

import (
	"log"
	"os"
	_ "search/matchers"
	"search/search"
)

func init() {
	log.Println("init..")
	log.SetOutput(os.Stdout)
}

func Main() {
	log.Println("starting search...")
	search.Run()
}
