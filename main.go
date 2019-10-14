package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var w Writer = ConsoleWriter{}
	n, err := w.Write([]byte("Hello Go"))
	fmt.Println(n, err)
	Channel()
}

// Writer @args []byte @return int, error
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter struct
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
