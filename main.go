package main

import (
	"fmt"
)

func main() {
	fmt.Println(Getenv("GITHUB_EVENT_PATH"))
}
