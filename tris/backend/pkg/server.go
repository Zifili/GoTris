package main

import (
	"GoTris/pkg/internal/server"
)

func main() {
	println("Starting listening on port 8080")
	srv := server.NewHTTPServer(":8080")
	srv.ListenAndServe()
}
