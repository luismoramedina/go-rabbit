package main

import (
  "fmt"
  "log"

  "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
    panic(fmt.Sprintf("%s: %s", msg, err))
  }
}

//Send a message
func Send() {
   conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
   failOnError(err, "Failed to connect to RabbitMQ")
   defer conn.Close()

   ch, err := conn.Channel()
   failOnError(err, "Failed to open a channel")
   defer ch.Close()

   q, err := ch.QueueDeclare(
     "hello", // name
     false,   // durable
     false,   // delete when unused
     false,   // exclusive
     false,   // no-wait
     nil,     // arguments
   )
   failOnError(err, "Failed to declare a queue")

   body := "hello from go"
   var headers = amqp.Table{}
   headers["spanTraceId"] = "go id"

   err = ch.Publish(
     "",     // exchange
     q.Name, // routing key
     false,  // mandatory
     false,  // immediate
     amqp.Publishing {
       Headers: headers,
       ContentType: "text/plain",
       Body:        []byte(body),
     })
   failOnError(err, "Failed to publish a message")
}