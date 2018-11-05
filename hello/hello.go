package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	var s = "Hello, world!"
	fmt.Println(s)
}
