package main

import (
	"log"
	"os"

	"github.com/jailtonjunior94/udemy-go-cli/app"
)

func main() {
	aplicacao := app.Gerar()
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
