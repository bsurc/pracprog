package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return ""
	}
	return fmt.Sprintf("Hello %s!", name)
}

func main() {
	var s = "world!"
	fmt.Println(hello(s))
}