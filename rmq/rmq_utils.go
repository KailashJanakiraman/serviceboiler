package rmq

import (
	"fmt"
	"log"
	"github.com/kailashjanakiraman/go-cloud/serviceboiler/db"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

//GetChannel ...
func GetChannel(url string)( * amqp.Connection,  * amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establoish connection to message broker")

	ch, err := conn.Channel()
	failOnError(err, "Failed to get a channel for the connection")
	return conn, ch
}

//InitExchange ...
func InitExchange(topicExchange string, ch * amqp.Channel) {
	err := ch.ExchangeDeclare(
		topicExchange, // name
		"topic", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)

	failOnError(err, "Failed to declare an exchange")
}

//GetQueue ...
func GetQueue(topicExchange string, ch * amqp.Channel) * amqp.Queue {
	InitExchange(topicExchange, ch)

	q, err := ch.QueueDeclare(
		"", // name
		false, // durable
		false, // delete when unused
		true, // exclusive
		false, // no-wait
		nil, // arguments
	)
	failOnError(err, "Failed to declare queue on the channel")
	return & q
}

//DecodeDBMessage ...
func DecodeDBMessage(theEncodedMessage []byte) * db.DatabaseTrigger {
	dbTrigger :=  & db.DatabaseTrigger {}	
	if  err := proto.Unmarshal(theEncodedMessage, dbTrigger); err != nil {
		log.Fatalln("Failed to parse Database Trigger:", err)
	}

	log.Printf("Received a message: Trigger=", dbTrigger.Trigger, ", Data1=", dbTrigger.Data1, ", Data2=", dbTrigger.Data2)
	return dbTrigger
}

//EncodeDBMessage ...
func EncodeDBMessage(data1 string, data2 []string)[]byte {
	log.Println("EncodeDBMessage(data1=", data1, ", data2=", data2)
	dbTrigger :=  & db.DatabaseTrigger {
		Trigger:db.DBTrigger_TRIGGER_1, 
		Data1:data1, 
		Data2:data2, 
	}

	msg, err := proto.Marshal(dbTrigger)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	return msg
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}

}
