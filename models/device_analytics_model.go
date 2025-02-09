package models

type Analytics struct {
	Step string `json:"stp"`
	Temp string `json:"temp"`
}

type AnalyticsResponse struct {
	CoOrdinates []*Analytics `json:"co_ordinates"`
}
