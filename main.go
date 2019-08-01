package main

import (
	"fmt"
	"os"
	"github.com/go-martini/martini"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "I was built and deployed by Concourse!"
	}

	m := martini.Classic()
	m.Get("/", func() string {
		fmt.Println(message)
		return message
	})
	m.Run()
}
