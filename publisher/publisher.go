package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker   = "ssl://37558b5eef6940f690902340b4a02081.s1.eu.hivemq.cloud:8883" // Replace with your broker URL and port
	topic    = "MYmessage"                                                      // Replace with your MQTT topic
	clientID = "mqtt-publisher"
)

func main() {
	// Set up MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername("sivan")     // Replace with your username if needed
	opts.SetPassword("Mqtt2024@") // Replace with your password if needed

	// Create and start an MQTT client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to the broker:", token.Error())
		os.Exit(1)
	}
	fmt.Println("Connected to MQTT broker")

	// Publish a message to the topic every 5 seconds
	for {
		text := fmt.Sprintf("Hello MQTT at %s", time.Now().Format(time.RFC3339))
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		fmt.Printf("Published message: %s\n", text)
		time.Sleep(5 * time.Second)
	}

	// Disconnect the client when done (this won't be reached in this loop)
	// client.Disconnect(250)
}
