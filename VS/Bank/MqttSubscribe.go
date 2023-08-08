package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

var client mqtt.Client
var id string
var responseChannel chan string

func StartMQTT() {
	go zeigMassege()
	go check()
	brokerURL := "tcp://localhost:1883"
	hostname := os.Getenv("BROKER_URL")
	if hostname != "" {
		brokerURL = "tcp://" + hostname + ":1883"
	}
	//uuid := uuid.New().String()
	id = os.Getenv("BANKNAME")
	if id == "" {
		id = "default"
	}
	opts := mqtt.NewClientOptions().AddBroker(brokerURL).SetClientID(id)
	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	responseChannel = make(chan string)
	Subscribe(client, responseChannel)

	//Publish(client)
}
func Publish(client mqtt.Client) {

	message1 := "Nachricht 1"
	token1 := client.Publish(Topic, 0, false, message1)
	token1.Wait()
	message2 := "Nachricht 2"
	token2 := client.Publish(Topic, 0, false, message2)
	token2.Wait()
}
func Subscribe(client mqtt.Client, responseChannel chan<- string) {

	if token := client.Subscribe(Topic, 1, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	client.Subscribe(Topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		//fmt.Printf("Empfangene Nachricht: %s von Thema: %s\n", msg.Payload(), msg.Topic())
		responseChannel <- string(msg.Payload())
	})
}
func zeigMassege() {
	BankHlp := ""
	commitCounter := 0
	for {
		select {
		case msg := <-responseChannel:
			fmt.Println("Nachricht empfangen:", msg) //Help!-BankName //commit //rollback
			if msg[:5] == "Help!" {
				BankHlp = msg[6:]
				//if BankHlp != id {
				//	performTransaction(msg[6:])
				//}
				if Inventor > 50000 {
					message1 := "commit"
					token1 := client.Publish(Topic, 0, false, message1)
					token1.Wait()
				} else {
					message1 := "Kann leider nicht helfen"
					token1 := client.Publish(Topic, 0, false, message1)
					token1.Wait()
				}
			} else if msg == "commit" {
				commitCounter++
				if commitCounter == 3 {
					if BankHlp != id {
						performTransaction(BankHlp)
					}
					commitCounter = 0
				}
			}
		default:
			// Tue nichts und warte weiter
		}

	}
}
func performTransaction(bankname string) {
	Inventor -= 100000
	RPCuberweisung("100000", bankname)

	//RPCuberweisung("100000", "localhost")

	//fmt.Println("Transaktion durchgeführt")
}
func sendeHelp(client mqtt.Client, id string) {
	payload := []byte("Help!-" + id)
	start = time.Now()
	if token := client.Publish(Topic, 1, false, payload); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error(), "\tcannt send Help")
	}
	fmt.Println("Help Anfrage gesendet")
}
func check() {
	for {
		if Inventor <= 0 {
			sendeHelp(client, id)
			time.Sleep(time.Second * 5)
		}
		//time.Sleep(time.Second)
	}
}
