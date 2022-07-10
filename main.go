package main

import (
	"ip-resolver/app"
	"os"
	"log"
)

func main() {
	application := app.Build()
	if err := application.Run(os.Args) ; err != nil {
		log.Fatal(err)
	}

}