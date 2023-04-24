package pkg

import "fmt"

type option func(*Foo) option

type Foo struct {
	verbosity int
}

// Option sets the options specified.
// It returns a slice of option to restore the corresponding arg's previous value.
func (f *Foo) Option(opts ...option) (previous []option) {
	previous = make([]option, len(opts))

	for i, opt := range opts {
		previous[i] = opt(f)
	}

	return previous
}

// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
	return func(f *Foo) option {
		previous := f.verbosity
		f.verbosity = v
		return Verbosity(previous)
	}
}

func (f *Foo) DoSomeDebugging() {
	fmt.Printf("verbosity: %d\n", f.verbosity)
}
