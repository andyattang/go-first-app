package myfunc

import "fmt"

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(str string) (n int, err error) {
	fmt.Println("exec in myfunc.Println.")
	return fmt.Println(str)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(str string, args ...interface{}) (n int, err error) {
	fmt.Println("exec in myfunc.Printf.")
	return fmt.Printf(str, args...)
}
