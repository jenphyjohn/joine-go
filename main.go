package main

import (
	"fmt"
	"github.com/jenphyjohn/joine-go/config"
)

func main() {
	value := config.GetValue("server.contextPath")
	fmt.Printf(value)
}
