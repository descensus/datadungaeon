package models

import "gorm.io/gorm"

type AqaraPlug struct {
	gorm.Model
	Name              string  `json:"name"`
	DeviceTemperature int     `json:"device_temperature"`
	Energy            float64 `json:"energy"`
	Linkquality       int     `json:"linkquality"`
	Power             float64 `json:"power"`
	State             string  `json:"state"`
}

type AqaraTemperature struct {
	gorm.Model
	Name             string  `json:"name"`
	Battery          int     `json:"battery"`
	Humidity         float64 `json:"humidity"`
	Linkquality      int     `json:"linkquality"`
	PowerOutageCount int     `json:"power_outage_count"`
	Pressure         float64 `json:"pressure"`
	Temperature      float64 `json:"temperature"`
	Voltage          int     `json:"voltage"`
}

type AqaraMagnet struct {
	gorm.Model
	Name              string `json:"name"`
	Battery           int    `json:"battery"`
	Contact           bool   `json:"contact"`
	DeviceTemperature int    `json:"device_temperature"`
	Linkquality       int    `json:"linkquality"`
	PowerOutageCount  int    `json:"power_outage_count"`
	Voltage           int    `json:"voltage"`
}

type Pocsag struct {
	gorm.Model
	Protocol string
	Address  string
	Message  string
}
