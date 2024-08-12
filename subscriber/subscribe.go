package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/mattn/go-sqlite3"
)

const (
	broker   = "ssl://37558b5eef6940f690902340b4a02081.s1.eu.hivemq.cloud:8883" // Replace with your broker URL and port
	topic    = "MYmessage"                 // Replace with your MQTT topic
	clientID = "mqtt-subscriber"
	username = "sivan"              // Replace with your username if required
	password = "Mqtt2024@"              // Replace with your password if required
)

func main() {
	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./mqtt_messages.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"topic" TEXT,
		"payload" TEXT,
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Set up MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username) // Include this line if your broker requires a username
	opts.SetPassword(password) // Include this line if your broker requires a password
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		// Print received message
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

		// Insert message into SQLite database
		insertSQL := `INSERT INTO messages(topic, payload) VALUES (?, ?)`
		_, err := db.Exec(insertSQL, msg.Topic(), string(msg.Payload()))
		if err != nil {
			log.Printf("Failed to insert message: %v", err)
		}
	})

	// Create and start an MQTT client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	// Subscribe to the topic
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Wait for interrupt signal to gracefully shutdown the subscriber
	sigChan := make(chan os.Signal, 1)
	<-sigChan

	// Disconnect the client and clean up
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}
