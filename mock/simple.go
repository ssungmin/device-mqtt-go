package main


import (
	"fmt"
	"os"
	"log"
	"time"
	"github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
  }

func onMessageReceived(client mqtt.Client , msg mqtt.Message) {
	fmt.Printf("ssang: %s\n", msg.Topic())
	fmt.Printf("ssangmsg: %s\n", msg.Payload()) }

func main2() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("hahaha")
	opts.SetUsername("admin")
	opts.SetPassword("public")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetDefaultPublishHandler(f)


	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("DataTopic", 0, onMessageReceived); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("DataTopic", 0, false, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)
	if token := c.Unsubscribe("DataTopic"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	c.Disconnect(250)
	time.Sleep(1 * time.Second)




}