package mqttcontroller

import (
	"encoding/json"
	"log"

	"github.com/Magowtham/dehydrater-server/repository"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttMessageHandler struct {
	repo repository.PostgresRepository
}

func NewMqttMessageHandler(repo repository.PostgresRepository) *MqttMessageHandler {
	return &MqttMessageHandler{
		repo: repo,
	}
}

func (h *MqttMessageHandler) HandleMqttMessage(c mqtt.Client, m mqtt.Message) {

	stepResponse, err := h.repo.GetSteps()

	if err != nil {
		log.Println(err)
		return
	}

	jsonResponse, err := json.Marshal(stepResponse)

	if err != nil {
		log.Println(err)
		return
	}

	c.Publish("d1", 1, false, jsonResponse)
}

func (h *MqttMessageHandler) HandleProcessStopMessage(c mqtt.Client, m mqtt.Message) {
	if err := h.repo.DeleteSteps(); err != nil {
		log.Println(err)
		return
	}
}
