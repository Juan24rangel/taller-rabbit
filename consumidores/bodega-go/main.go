package main

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	host := os.Getenv("RABBIT_HOST")
	if host == "" {
		host = "localhost"
	}

	conn, err := amqp.Dial("amqp://guest:guest@" + host + ":5672/")
	if err != nil {
		log.Fatalf("Error conectando: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error abriendo canal: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("eventos", "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declarando exchange: %v", err)
	}

	_, err = ch.QueueDeclare("cola_bodega", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declarando cola: %v", err)
	}

	err = ch.QueueBind("cola_bodega", "inventario.bajo", "eventos", false, nil)
	if err != nil {
		log.Fatalf("Error en binding: %v", err)
	}

	mensajes, err := ch.Consume("cola_bodega", "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error consumiendo: %v", err)
	}

	fmt.Println("[Bodega] Esperando mensajes... (Ctrl+C para salir)")

	for m := range mensajes {
		fmt.Println("[Bodega] Recibido:", string(m.Body))
		m.Ack(false)
	}
}
