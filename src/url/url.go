package url

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Main() {
	for _, url := range os.Args[1:] {
		checkUrl(url)
	}
}

func checkUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch loading body: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s", b)
}
