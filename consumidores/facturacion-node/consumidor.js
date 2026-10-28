const amqp = require("amqplib");

const RABBIT_HOST = process.env.RABBIT_HOST || "localhost";

async function main() {
  const conexion = await amqp.connect(`amqp://${RABBIT_HOST}`);
  const canal = await conexion.createChannel();

  await canal.assertExchange("eventos", "topic", { durable: true });
  await canal.assertQueue("cola_facturacion", { durable: true });
  await canal.bindQueue("cola_facturacion", "eventos", "pago.aprobado");

  console.log("[Facturacion] Esperando mensajes... (Ctrl+C para salir)");

  canal.consume("cola_facturacion", (msg) => {
    if (msg !== null) {
      const mensaje = JSON.parse(msg.content.toString());
      console.log("[Facturacion] Recibido:", mensaje);
      canal.ack(msg);
    }
  });
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
