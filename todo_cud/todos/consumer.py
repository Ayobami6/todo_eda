import logging
from kafka import KafkaConsumer
from .tododb_service import TodoDBService

logger = logging.getLogger(__name__)


class TodoConsumer:
    """Consumer class for processing todo messages."""

    def __init__(self):
        self.consumer = KafkaConsumer(
            "todos",
            bootstrap_servers="localhost:9092",
            group_id="todo_group",
            auto_offset_reset="earliest",
            enable_auto_commit=True,
            value_deserializer=lambda x: x.decode("utf-8"),
            key_deserializer=lambda k: k.decode("utf-8") if k else None,
        )

    def consume_messages(self):
        """Consume messages from the Kafka topic."""
        try:

            for message in self.consumer:
                logger.info(f"Received message: {message.value}")
                result = TodoDBService.process_message(message.value)
                if result:
                    self.consumer.commit()
                    logger.info(f"Processed message successfully: {result}")
                else:
                    logger.error(f"Failed to process message: {message.value}")

        except KeyboardInterrupt:
            logger.info("Stopping consumer...")
        except Exception as e:
            logger.error(f"Error consuming messages: {e}")
        finally:
            self.consumer.close()
            logger.info("Consumer closed.")
