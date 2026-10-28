import pika
import json
import time
import os
from datetime import datetime

RABBIT_HOST = os.getenv("RABBIT_HOST", "localhost")

# Conexión al broker
conexion = pika.BlockingConnection(pika.ConnectionParameters(host=RABBIT_HOST))
canal = conexion.channel()

# Declarar el exchange tipo topic
canal.exchange_declare(exchange="eventos", exchange_type="topic", durable=True)

contador = 1
while True:
    mensaje = {
        "id": contador,
        "evento": "pedido.creado",
        "origen": "productor-pedidos-python",
        "hora": datetime.now().isoformat()
    }
    canal.basic_publish(
        exchange="eventos",
        routing_key="pedido.creado",
        body=json.dumps(mensaje)
    )
    print(f"[Pedidos] Enviado: {mensaje}", flush=True)
    contador += 1
    time.sleep(3)
