package models

type DeviceStatus struct {
	PresentStep        int32   `json:"stp"`
	PresentTemperature float64 `json:"tmp"`
}
