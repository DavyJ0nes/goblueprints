package trace

import (
	"bytes"
	"testing"
)

// TestNew checks that New() can create new Tracer object
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace Error: Got: %s | Expected: Hello trace package", buf.String())
		}
	}
}

// TestOff Checks that a blank or nilTracer can be created
func TestOff(t *testing.T) {
	var buf bytes.Buffer
	var silentTracer Tracer = Off()
	silentTracer.Trace("something")
	if buf.String() != "" {
		t.Errorf("Trace Error. Got: %s | Expected: ''", buf.String())
	}
}
