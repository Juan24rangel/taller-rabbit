import pika
import json
import os

RABBIT_HOST = os.getenv("RABBIT_HOST", "localhost")

conexion = pika.BlockingConnection(pika.ConnectionParameters(host=RABBIT_HOST))
canal = conexion.channel()

# 1. Declarar el exchange (igual que el productor, por si arranca primero el consumidor)
canal.exchange_declare(exchange="eventos", exchange_type="topic", durable=True)

# 2. Declarar la cola
canal.queue_declare(queue="cola_notificaciones", durable=True)

# 3. Binding: conectar la cola al exchange con la regla "pedido.creado"
canal.queue_bind(exchange="eventos", queue="cola_notificaciones", routing_key="pedido.creado")

def procesar(ch, method, properties, body):
    mensaje = json.loads(body)
    print(f"[Notificaciones] Recibido: {mensaje}", flush=True)
    ch.basic_ack(delivery_tag=method.delivery_tag)

# 4. Empezar a consumir
canal.basic_consume(queue="cola_notificaciones", on_message_callback=procesar)
print("[Notificaciones] Esperando mensajes... (Ctrl+C para salir)", flush=True)
canal.start_consuming()
