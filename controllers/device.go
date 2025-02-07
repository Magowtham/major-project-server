package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Magowtham/dehydrater-server/models"
)

type DeviceHandler struct {
	repo models.DeviceRepo
}

func NewDeviceHandler(repo models.DeviceRepo) *DeviceHandler {
	return &DeviceHandler{
		repo: repo,
	}
}

func (h *DeviceHandler) DeviceAddStepHandler(w http.ResponseWriter, r *http.Request) {
	var deviceStepRequest models.DeviceStepRequest
	if err := json.NewDecoder(r.Body).Decode(&deviceStepRequest); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"messsage": "invalid request body format"})
	}

	if err := h.repo.AddStep(deviceStepRequest.Steps); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		json.NewEncoder(w).Encode(err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "steps added successfully"})

}

func (h *DeviceHandler) GetStepsHandler(w http.ResponseWriter, r *http.Request) {
	steps, err := h.repo.GetSteps()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(steps)
}
