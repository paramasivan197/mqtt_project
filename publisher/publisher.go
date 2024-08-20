package main

import (
	"fmt"
	"os"

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
		// Read message from the user
		fmt.Print("Enter message to publish (or type 'exit' to quit): ")
		var text string
		fmt.Scanln(&text)
	
		// Exit the loop if the user types "exit"
		if text == "exit" {
			break
		}
	
		// Publish the user-provided message to the MQTT topic
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		fmt.Printf("Published message: %s\n", text)
		
		// Optional: Add a delay if needed
		// time.Sleep(5 * time.Second)
	}
	
	// Disconnect the client when done
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
	
	
}
