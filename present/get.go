package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// START OMIT
	// Go (ignoring errors for brevity)
	var resp *http.Response
	resp, _ = http.Get("https://example.com/")
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	// END OMIT
}
