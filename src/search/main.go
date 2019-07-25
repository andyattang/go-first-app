package search

import (
	"log"
	"os"
	"search/search"
)

func init() {
	log.Println("init..")
	log.SetOutput(os.Stdout)
}

func Main() {
	log.Println("starting search...")
	search.Run("president")
}
