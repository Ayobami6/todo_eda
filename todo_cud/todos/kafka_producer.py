from django.conf import settings
from kafka import KafkaProducer
from kafka.errors import KafkaError
import json
import logging
from datetime import datetime


logger = logging.getLogger(__name__)


class TodoKafkaProducer:
    def __init__(self):
        self.producer = KafkaProducer(
            bootstrap_servers=settings.KAFKA_BROKER_URL,
            value_serializer=lambda v: json.dumps(v).encode("utf-8"),
            retries=5,
            key_serializer=lambda k: k.encode("utf-8") if k else None,
            acks="all",  # Wait for the leader to acknowledge the record
            enable_idempotence=True,  # Ensure no duplicate messages
        )

    def send_message(self, data, topic, ops):
        message = {
            "operation": ops,
            "data": data,
            "timestamp": datetime.utcnow().isoformat(),
        }
        try:
            future = self.producer.send(topic, message)
            result = future.get(
                timeout=10
            )  # Block until a single message is sent (or timeout)
            logger.info(f"Message sent to topic {topic}: {message}")
            return result
        except KafkaError as e:
            logger.error(f"Failed to send message to topic {self.topic}: {e}")
            return None

    def close(self):
        self.producer.flush()
        self.producer.close()


todo_producer = TodoKafkaProducer()

__all__ = ["todo_producer"]
