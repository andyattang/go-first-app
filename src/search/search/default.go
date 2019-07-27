package search

import (
	"fmt"
	"log"
)

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	log.Println("default search: ", feed)
	return nil, nil
}

func Match(matcher Matcher, feed *Feed, results chan<- *Result) {
	var searchTerm = feed.Search
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("URL:%s\n%s:%s\n\n", result.Url, result.Field, result.Search)
	}
}
