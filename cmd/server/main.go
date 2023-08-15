package main

import "ddd/application/server"

func main() {
	app := server.New()

	app.Start()
}
