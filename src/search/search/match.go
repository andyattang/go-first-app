package search

type Result struct {
	Url     string
	Field   string
	Content string
	Search  string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}
