package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Magowtham/dehydrater-server/db"
	mqttcontroller "github.com/Magowtham/dehydrater-server/mqtt-controller"
	"github.com/Magowtham/dehydrater-server/repository"
	"github.com/Magowtham/dehydrater-server/routes"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

	//mqtt

	mqttHandler := mqttcontroller.NewMqttMessageHandler(*postgresRepo)

	opts := mqtt.NewClientOptions()

	var brokerAddress = fmt.Sprintf("tcp://%s:%s", "34.47.250.228", "1883")

	opts.AddBroker(brokerAddress)
	opts.SetClientID("majorproject/message/processor")

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("connected to broker")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Println("disconnected from broker", err.Error())
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			log.Printf("error occurred while connecting to broker, Error -> %v\n", token.Error())
			return
		}

		c.Unsubscribe("/message/processor")
		c.Unsubscribe("/stop/process")
		c.Subscribe("/message/processor", 1, mqttHandler.HandleMqttMessage)
		c.Subscribe("/stop/process", 1, mqttHandler.HandleProcessStopMessage)
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("error occurred while connecting to broker, error -> %v\n", token.Error())
	}

	if client.IsConnected() {
		client.Subscribe("/message/processor", 1, mqttHandler.HandleMqttMessage)
		client.Subscribe("/stop/process", 1, mqttHandler.HandleProcessStopMessage)
	}

	route := routes.Router(postgresRepo)

	log.Printf("http server is listening on: %v\n", serverAddress)
	http.ListenAndServe(serverAddress, route)

}
