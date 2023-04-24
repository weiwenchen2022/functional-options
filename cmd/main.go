package main

import pkg "functional-option"

func main() {
	var foo pkg.Foo

	DoSomethingVerbosely(&foo, 3)
	foo.DoSomeDebugging()
}

func DoSomethingVerbosely(foo *pkg.Foo, verbosity int) {
	// Could combine the next two lines,
	// with some loss of readability.
	prev := foo.Option(pkg.Verbosity(verbosity))[0]
	defer foo.Option(prev)

	// ... do some stuff with foo under high verbosity.
	foo.DoSomeDebugging()
}
