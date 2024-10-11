package main

import (
	"one-million-checkboxes/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
