package mqttcontroller

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/Magowtham/dehydrater-server/models"
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

	if err := h.repo.DeleteAnalytics(); err != nil {
		log.Println(err)
		return
	}
}

func (h *MqttMessageHandler) HandleAnalyticsUpdate(c mqtt.Client, m mqtt.Message) {
	var step models.DeviceStatus

	if err := json.Unmarshal(m.Payload(), &step); err != nil {
		log.Println(err)
		return
	}

	presentStep := strconv.Itoa(int(step.PresentStep) + 1)

	presentTemp := strconv.Itoa(int(step.PresentTemperature))

	if err := h.repo.UpdateAnalytics(presentStep, presentTemp); err != nil {
		log.Println(err)
		return
	}
}
