package service

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"time"
)

func failError(err error,msg string){
	if err != nil{
		log.Fatalf("%s: %s", msg, err)
	}
}

func Order(m Message)error{
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failError(err, "Can't connect to MQ")
	if err!=nil {return err}
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	//failError(err, "Can't create a Channel")
	if err!=nil {return err}
	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("good",true,false,false,false,nil)
	//failError(err, "Could not declare queue")
	if err!=nil {return err}

	rand.Seed(time.Now().UnixNano())

	body, err:= json.Marshal(m)
	if err != nil{
		//failError(err, "Error encoding JSON")
		return err
	}

	err = amqpChannel.Publish("",queue.Name,false,false,amqp.Publishing{
		DeliveryMode : amqp.Persistent,
		ContentType : "text/plain",
		Body : body,
	})
	if err != nil{
		//log.Fatalf("Error publishing message: %s",err)
		return err
	}

	log.Printf("AddMessage: %s",string(body))
	return nil
}