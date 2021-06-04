package main

import (
	"errors"
	"fmt"
	"os"
)

type ErrorA struct{}

func (*ErrorA) Error() string {
	return "ErrorA"
}

type ErrorB struct{}

func (*ErrorB) Error() string {
	return "ErrorB"
}

func (*ErrorB) Unwrap() error {
	return os.ErrNotExist
}

func handle() (string, bool) {
	return "", false
}

func main() {
	noEmbedErr := fmt.Errorf("Failed bla bla")                  // no embedding
	embedErr := fmt.Errorf("Failed bla bla %w", os.ErrNotExist) // %w embeds error within error
	anotherEmbedErr := fmt.Errorf("Lol %w", embedErr)           // another level of embedding

	// print types
	fmt.Printf("%T\n", noEmbedErr) // *errors.errorString
	fmt.Printf("%T\n", embedErr)   // *fmt.wrapError

	fmt.Println(embedErr)                                   // prints `Failed bla bla file does not exist`, %w works the same as %v + embedding
	fmt.Println(errors.Is(embedErr, os.ErrNotExist))        // true
	fmt.Println(errors.Is(anotherEmbedErr, os.ErrNotExist)) // true, because it checks all errors in the chain

	fmt.Println(errors.Is(&ErrorA{}, os.ErrNotExist)) // false
	fmt.Println(errors.Is(&ErrorB{}, os.ErrNotExist)) // true, because it implements Unwrap

	// -----
	unknownErr := &ErrorB{} // let's assume we don't know the type

	var errB *ErrorB // we are testing against ErrorB
	if errors.As(unknownErr, &errB) {
		fmt.Println(errB) // ErrorB
	}
}
