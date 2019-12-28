package main

import "github.com/ahmadkarlam/go-message/pkg/api"

func main() {
	err := api.Start("8000")
	if err != nil {
		panic(err.Error())
	}
}
