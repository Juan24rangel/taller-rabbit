package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

	contador := 1
	for {
		mensaje := map[string]interface{}{
			"id":     contador,
			"evento": "inventario.bajo",
			"origen": "productor-inventario-go",
			"hora":   time.Now().Format(time.RFC3339),
		}
		cuerpo, _ := json.Marshal(mensaje)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = ch.PublishWithContext(ctx, "eventos", "inventario.bajo", false, false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        cuerpo,
			})
		cancel()
		if err != nil {
			log.Fatalf("Error publicando: %v", err)
		}

		fmt.Println("[Inventario] Enviado:", string(cuerpo))
		contador++
		time.Sleep(3 * time.Second)
	}
}
