package routes

import (
	"github.com/Magowtham/dehydrater-server/controllers"
	"github.com/Magowtham/dehydrater-server/models"
	"github.com/gorilla/mux"
)

func Router(repo models.DeviceRepo) *mux.Router {
	deviceHandler := controllers.NewDeviceHandler(repo)

	router := mux.NewRouter()

	router.HandleFunc("/steps", deviceHandler.DeviceAddStepHandler).Methods("POST")
	router.HandleFunc("/steps", deviceHandler.GetStepsHandler).Methods("GET")

	return router
}
