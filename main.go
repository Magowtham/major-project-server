package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Magowtham/dehydrater-server/db"
	"github.com/Magowtham/dehydrater-server/repository"
	"github.com/Magowtham/dehydrater-server/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		log.Fatalln("SERVER_ADDRESS varaible not found")
	}

	dbConn, err := db.Connect()

	log.Println("connected to database")

	if err != nil {
		log.Fatalln(err)
	}

	postgresRepo := repository.NewPostgresRepository(dbConn)

	if err := postgresRepo.Init(); err != nil {
		log.Fatalln(err)
	}

	log.Println("database initialized succesfully")

	route := routes.Router(postgresRepo)

	log.Printf("http server is listening on: %v\n", serverAddress)
	http.ListenAndServe(serverAddress, route)

}
