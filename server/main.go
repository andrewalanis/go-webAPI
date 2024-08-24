package main

import (
	"fmt"
	"gowebapi/initializers"
)



func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("Hello, World!")
}
