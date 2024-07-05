/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/gabrielAnFran/frauCLI/cmd"
)

func main() {
	err := godotenv.Load() 
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
