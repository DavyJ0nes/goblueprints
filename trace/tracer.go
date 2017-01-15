package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that desribes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

// nilTracer is blank to be used with Off()
type nilTracer struct{}

// New instanciates new Tracer object
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// Trace outputs formatted text to io.Writer
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// Trace is a nilTracer version of the function to satisy the Tracer interface
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to trace
func Off() Tracer {
	return &nilTracer{}
}
