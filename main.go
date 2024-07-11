package main

import (
	"log"
	"os"
	"simple-cli/app"
)

func main() {
	
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}