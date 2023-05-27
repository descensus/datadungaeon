package main

import (
	"datadungaeon/color"
	"datadungaeon/database"
	"datadungaeon/models"
	"fmt"
	"log"
	"strconv"
	"strings"

	//"time"
	"encoding/json"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// const appName = "datadungæon"
const appName = "datadungaeon"

var RandColor string
var (
	//c1 = color.RandColor()
	c2 = color.RandColor()
	c3 = color.RandColor()
	c4 = color.RandColor()
)
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	clopt := client.OptionsReader()
	log.Printf("(%s) * Connected to MQTT Broker: %s\n", appName, clopt.Servers()[0])
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	clopt := client.OptionsReader()
	log.Printf("(%s) * Connect lost to: %s (Error: %v)\n", appName, clopt.Servers()[0], err)
}

var msgAqaraPlug mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	msg_json := msg.Payload()
	var aqaraplug models.AqaraPlug
	err := json.Unmarshal(msg_json, &aqaraplug)
	if err != nil {
		fmt.Println("error:", err)
	}
	aqaraplug.Name = msg.Topic()
	database.Instance.Save(&aqaraplug)
	log.Printf("(%s) + %s[%s] Current Power: %.2f W Total: %.2f kWh%s\n", appName, RandColor, msg.Topic(), aqaraplug.Power, aqaraplug.Energy, color.Reset)

}

var msgAqaraMagnet mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	msg_json := msg.Payload()
	var aqaraMag models.AqaraMagnet
	err := json.Unmarshal(msg_json, &aqaraMag)
	if err != nil {
		fmt.Println("error:", err)
	}
	aqaraMag.Name = msg.Topic()
	database.Instance.Save(&aqaraMag)
	log.Printf("(%s) > %s[%s] Contact: %t%s\n", appName, c2, msg.Topic(), aqaraMag.Contact, color.Reset)

}

var msgAqaraTemp mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	msg_json := msg.Payload()
	var aqaraTemp models.AqaraTemperature
	err := json.Unmarshal(msg_json, &aqaraTemp)
	if err != nil {
		log.Println("error:", err)
	}
	aqaraTemp.Name = msg.Topic()
	database.Instance.Save(&aqaraTemp)
	log.Printf("(%s) * %s[%s] Current Temp: %.2fc Humidity: %.2f %% %s\n", appName, c3, msg.Topic(), aqaraTemp.Temperature, aqaraTemp.Humidity, color.Reset)

}

var msgPocsag mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var pocsag models.Pocsag
	text := string(msg.Payload())

	result := strings.Split(text, ",")

	protocol := result[0]
	address := result[1]
	message := result[2:]

	log.Printf("(%s) # %s[%s] Address: %s Protocol: %s Message: %s\n%s", appName, c4, msg.Topic(), address, protocol, message, color.Reset)

	pocsag.Address = address
	pocsag.Protocol = protocol
	pocsag.Message = strings.Join(message, " ")
	database.Instance.Save(&pocsag)

}

func main() {

	var host string = os.Getenv("DATADUNGAEON_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DATADUNGAEON_DB_PORT"))
	if err != nil {
		log.Fatalln("DB Port is not an integer")
	}
	var username string = os.Getenv("DATADUNGAEON_DB_USERNAME")
	var password string = os.Getenv("DATADUNGAEON_DB_PASSWORD")
	var databaseName string = os.Getenv("DATADUNGAEON_DB_DBNAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, databaseName)
	database.Connect(connectionString)
	database.Migrate()

	var clientID string = os.Getenv("DATADUNGAEON_MQTT_CLIENTID")
	var mqtt_broker string = os.Getenv("DATADUNGAEON_MQTT_HOST")
	mqtt_port, err := strconv.Atoi(os.Getenv("DATADUNGAEON_MQTT_PORT"))
	if err != nil {
		log.Fatalln("MQTT Port is not an integer")
	}
	var mqtt_username string = os.Getenv("DATADUNGAEON_MQTT_USERNAME")
	var mqtt_password string = os.Getenv("DATADUNGAEON_MQTT_PASSWORD")
	c := make(chan os.Signal, 1)
	opts := mqtt.NewClientOptions()
	broker_connstring := fmt.Sprintf("tcp://%s:%d", mqtt_broker, mqtt_port)
	opts.AddBroker(broker_connstring)
	opts.SetClientID(clientID)
	opts.SetUsername(mqtt_username)
	opts.SetPassword(mqtt_password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	/*
	   Cycle through all the environment variablees so we know which sensors to subscribe to.
	*/
	devs := [...]string{"AQARA_PLUG", "AQARA_TEMPERATURE", "AQARA_MAGNET", "POCSAG"}

	for _, device := range devs {

		i := 0

	OKBAI:
		for {
			key := fmt.Sprintf("%s%d", device, i)
			value, exists := os.LookupEnv(key)
			if !exists {
				break OKBAI
			}
			if device == "AQARA_PLUG" {

				go aqaraPlug(client, value)

			} else if device == "AQARA_TEMPERATURE" {

				go aqaraTemp(client, value)

			} else if device == "AQARA_MAGNET" {

				go aqaraMagnet(client, value)

			} else if device == "POCSAG" {

				go pocsag(client, value)

			}
			i++
		}

	}

	<-c

}

func aqaraPlug(client mqtt.Client, topic string) {
	RandColor = color.RandColor()
	c := make(chan os.Signal, 1)
	client.AddRoute(topic, msgAqaraPlug)
	token := client.Subscribe(topic, 1, nil)
	log.Printf("Subscribed to topic: %s\n", topic)
	token.Wait()
	<-c
}

func aqaraTemp(client mqtt.Client, topic string) {
	c := make(chan os.Signal, 1)
	client.AddRoute(topic, msgAqaraTemp)
	token := client.Subscribe(topic, 1, nil)
	log.Printf("Subscribed to topic: %s\n", topic)
	token.Wait()
	<-c
}

func aqaraMagnet(client mqtt.Client, topic string) {
	c := make(chan os.Signal, 1)
	client.AddRoute(topic, msgAqaraMagnet)
	token := client.Subscribe(topic, 1, nil)
	log.Printf("Subscribed to topic: %s\n", topic)
	token.Wait()
	<-c
}

func pocsag(client mqtt.Client, topic string) {
	c := make(chan os.Signal, 1)
	client.AddRoute(topic, msgPocsag)
	token := client.Subscribe(topic, 1, nil)
	log.Printf("Subscribed to topic: %s\n", topic)
	token.Wait()
	<-c
}
