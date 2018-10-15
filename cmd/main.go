package main

import (
	"github.com/AnaTofuZ/AutoEBL"
	"log"
	"os"
)

func main() {
	autoebl := AutoEBL.New()
	if err := autoebl.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
