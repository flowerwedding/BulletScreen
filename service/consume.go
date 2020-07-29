package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func handleError(err error,msg string){
	if err != nil{
		log.Fatalf("%s:%s",msg,err)
	}
}

func OpenConsumer(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err,"Can't connect to MQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err,"Can't create a ampqChannel")
	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("good",true,false,false,false,nil)
	handleError(err, "Could not declare 'add' queue")

    err = amqpChannel.Qos(1,0,false)
	handleError(err,"Could not configue QoS")

    messageChannel, err := amqpChannel.Consume(queue.Name,"",false,false,false,false,nil)
	handleError(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func(){
	    log.Printf("Consumer ready,PID: %d",os.Getpid())
		for d := range messageChannel{
			log.Printf("Reeived a message: %s",string(d.Body))

			message := &Message{}
			err := json.Unmarshal(d.Body, message)
			if err != nil {
				log.Printf("Error decoding JSON: %s",err)
			}
			log.Printf("Message: %s",string(d.Body))

            if err := AddMessage(*message);err != nil{
				log.Printf("Error of storing mysql: %s",err)
			}

			if err := d.Ack(false); err != nil{
				log.Printf("Error acknowledging message : %s",err)
			}else{
				log.Printf("Acknowledeged message")
			}
		}
	}()

	<-stopChan
}
