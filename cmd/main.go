package main

import interfaces "github.com/muriloGervasio/golang-ddd-account/internal/interface"

func main() {
	server := interfaces.NewServer()

	server.Run()
}
