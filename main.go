package main

import (
	"flag"
	"gp/backend/app"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()
	app.Run(*addr)
}
