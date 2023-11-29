package main

import (
	"fmt"
	"log"
	"os"

	"github.com/582-21W-MA/barbershop/cmd/barbershop"
	"github.com/582-21W-MA/barbershop/cmd/serve"
	"github.com/582-21W-MA/barbershop/cmd/watch"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(usage)
		os.Exit(0)

	}
	if os.Args[1] == "serve" {
		if len(os.Args) < 3 {
			log.Fatalln("Argument <root_directory> missing. See usage.")
		}
		rootDir := os.Args[2]
		serve.Run(rootDir)
	}
	if os.Args[1] == "watch" {
		rootDir := os.Args[2]
		watch.Run(rootDir)
	}
	inputDir := os.Args[1]
	if err := barbershop.Run(inputDir); err != nil {
		log.Fatalf("Error %v", err)
	}
}
