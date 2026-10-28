FROM python:3.12-slim
WORKDIR /app
RUN pip install --no-cache-dir pika
COPY consumidor.py .
CMD ["python", "consumidor.py"]
