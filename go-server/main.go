package main

import (
	"one-million-checkboxes/routes"
)

// Package classification Account API.
//
// this is to show how to write RESTful APIs in golang.
// that is to provide a detailed overview of the language specs
//
// Terms Of Service:
//
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Supun Muthutantri<mydocs@example.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//
// swagger:meta

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
