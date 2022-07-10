package main

import (
	"fmt"
	"command-line/app"
	"os"
	"log"
)

func main() {
	fmt.Println("Entry point")

	application := app.Build()
	if err := application.Run(os.Args) ; err != nil {
		log.Fatal(err)
	}

}