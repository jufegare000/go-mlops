package main

import (
	"fmt"
	"ports/main/httpadapter"
)

func main() {
	httpadapter.RestRouter()
	fmt.Println("Welcome to this application on")
}
