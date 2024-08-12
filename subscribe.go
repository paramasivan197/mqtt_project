package main

import (
    "database/sql"
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func main() {
    // MQTT setup
    opts := mqtt.NewClientOptions().AddBroker("ssl://37558b5eef6940f690902340b4a02081.s1.eu.hivemq.cloud:8883") // Replace with your broker URL and port
    opts.SetClientID("go_mqtt_subscriber")

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    // SQLite setup
    db, err := sql.Open("sqlite3", "./mqtt_data.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the table if it doesn't exist
    createTableSQL := `CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        topic TEXT,
        message TEXT,
        received_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
    
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    // Subscribe and store messages
    client.Subscribe("test/topic", 0, func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

        insertSQL := `INSERT INTO messages (topic, message) VALUES (?, ?)`
        _, err := db.Exec(insertSQL, msg.Topic(), msg.Payload())
        if err != nil {
            fmt.Println("Error storing message:", err)
        }
    })

    // Keep the subscriber running
    select {}
}
