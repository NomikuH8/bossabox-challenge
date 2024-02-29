package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	file, err := os.ReadFile(path.Join(".", ".example.env"))
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile(".env", file, 0644)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Criado env com sucesso")
}
