package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in environment.")
	}
	

	fmt.Println("Port:", portString)
}