package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "time"
)

func main() {
    opts := mqtt.NewClientOptions().AddBroker("ssl://37558b5eef6940f690902340b4a02081.s1.eu.hivemq.cloud:8883")
    opts.SetClientID("go_mqtt_publisher")
    opts.SetUsername("sivan")
    opts.SetPassword("Mqtt2024@")

    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    topic := "occasion"
    message := "Hello, MQTT"
    token := client.Publish(topic, 1, false, message)
    token.Wait()

    fmt.Println("Published message:", message)
    time.Sleep(1 * time.Second)

    client.Disconnect(250)
}
                                      