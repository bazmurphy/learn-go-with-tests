package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// We want to write a function that greets someone,
// just like we did in the hello chapter
// but this time we are going to be testing the actual printing.

// func Greet(name string) {
// 	fmt.Printf("Hello, %s", name)
// }

// But how can we test this?
// Calling fmt.Printf prints to stdout,
// which is pretty hard for us to capture using the testing framework.

// What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing.

// Our function doesn't need to care where or how the printing happens,
// so we should accept an interface rather than a concrete type.

// If we do that, we can then change the implementation to print to something we control so that we can test it.
// In "real life" you would inject in something that writes to stdout.

// If you look at the source code of `fmt.Printf` you can see a way for us to hook in

// It returns the number of bytes written and any write error encountered.
// func Printf(format string, a ...interface{}) (n int, err error) {
// 	return Fprintf(os.Stdout, format, a...)
// }

// Interesting! Under the hood Printf just calls Fprintf passing in os.Stdout.

// What exactly is an os.Stdout? What does Fprintf expect to get passed to it for the 1st argument?

// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
// 	p := newPrinter()
// 	p.doPrintf(format, a)
// 	n, err = w.Write(p.buf)
// 	p.free()
// 	return
// }

// An io.Writer

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// From this we can infer that os.Stdout implements io.Writer;
// Printf passes os.Stdout to Fprintf which expects an io.Writer.

// As you write more Go code you will find this interface popping up a lot
// because it's a great general purpose interface for "put this data somewhere".

// So we know under the covers we're ultimately using Writer to send our greeting somewhere.
// Let's use this existing abstraction to make our code testable and more reusable.

// func Greet(writer *bytes.Buffer, name string) {
// 	// [0]
// 	// fmt.Printf("Hello, %s", name)

// 	// [1] Use the writer to send the greeting to the buffer in our test.
// 	// Remember fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to,
// 	// whereas fmt.Printf defaults to stdout.
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// Earlier the compiler told us to pass in a pointer to a bytes.Buffer. This is technically correct but not very useful.
// To demonstrate this, try wiring up the Greet function into a Go application where we want it to print to stdout.

// "cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet"

// As discussed earlier fmt.Fprintf allows you to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.

// [3] If we change our code to use the more general purpose interface
// we can now use it in both tests and in our application.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// The Internet

// [4] When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request that was used to make the request.
// When you implement your server you write your response using the writer.
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// [2] try to call the Greet function with os.Stdout
	Greet(os.Stdout, "Baz")
	// cannot use os.Stdout (variable of type *os.File) as *bytes.Buffer value in argument to GreetcompilerIncompatibleAssign

	// [5]
	// fmt.Fprintf allows you to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
