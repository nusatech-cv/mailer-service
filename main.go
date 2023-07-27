package main

import (
	"mailer-service/mailer"
	"mailer-service/rabbitmq"
	"mailer-service/token"
)

func main() {
	rabbitmq.Receive(func(record *token.Record) {
		mailer.SendMail(record)
	})
}
