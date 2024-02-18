package main

import "github.com/muriloGervasio/golang-ddd-account/internal/infrastructure"

func main() {
	server := infrastructure.NewServer()

	server.Run()
}
