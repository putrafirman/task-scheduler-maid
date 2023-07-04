package main

import (
	"log"

	"github.com/joho/godotenv"
	hybrid_serverless "putrafirman.com/playground/task-scheduler-maid"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to read .env", err.Error())
	}
	e := hybrid_serverless.InitRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
