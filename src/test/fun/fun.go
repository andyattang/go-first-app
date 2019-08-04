package fun

import (
	"fmt"
)

type myfunc func(int, string) string

func a(x int, s string) string {
	return fmt.Sprintln(x, s)
}

func test(f myfunc, x int, s string) {
	s1 := f(x, s)
	fmt.Print(s1)
}

func Main() {
	x := 1
	s := "string"
	test(a, x, s)
}
