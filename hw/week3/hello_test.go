package main

import (
	"fmt"
	"strings"
	"testing"
)

// HelloWorld returns the famous 'Hello, World!' string, substituting 'World'
// with the name provided by the caller.  If name is an empty string, or a
// string that contains only white space, an empty string is returned.
func HelloWorld(name string) string {
	if strings.TrimSpace(name) == "" {
		return ""
	}
	s := fmt.Sprintf("Hello, %s!", name)
	return s
}

// Write tests for the function above.  Identify special cases, and make sure
// you test for them.  At a minumum, two cases should be tested, a third may be
// more clear, but not necessary.  Make sure to read the comment to understand
// what the function *should* do.
func TestHelloWorld(t *testing.T) {
	// Technically, if you delete this line, the test passes.  That won't fly.
	t.Fatal("not implemented")
}
