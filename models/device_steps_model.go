package models

type DeviceStep struct {
	StepNumber  string `json:"step_number"`
	Time        string `json:"time"`
	Temperature string `json:"temp"`
}

type DeviceStepRequest struct {
	Steps []*DeviceStep `json:"steps"`
}

type DeviceStepResponse struct {
	Steps []*DeviceStep `json:"steps"`
}

type DeviceRepo interface {
	AddStep(steps []*DeviceStep) error
	GetSteps() (*DeviceStepResponse, error)
}
