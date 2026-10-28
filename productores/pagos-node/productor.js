const amqp = require("amqplib");

const RABBIT_HOST = process.env.RABBIT_HOST || "localhost";

async function main() {
  const conexion = await amqp.connect(`amqp://${RABBIT_HOST}`);
  const canal = await conexion.createChannel();

  await canal.assertExchange("eventos", "topic", { durable: true });

  let contador = 1;
  setInterval(() => {
    const mensaje = {
      id: contador,
      evento: "pago.aprobado",
      origen: "productor-pagos-node",
      hora: new Date().toISOString(),
    };
    canal.publish("eventos", "pago.aprobado", Buffer.from(JSON.stringify(mensaje)));
    console.log("[Pagos] Enviado:", mensaje);
    contador++;
  }, 3000);
}

main().catch((err) => {
  console.error("Error:", err);
  process.exit(1);
});
