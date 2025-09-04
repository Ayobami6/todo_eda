"""Contains services for creating, updating, and deleting todo items using Kafka for messaging."""

from .kafka_producer import todo_producer
import logging

logger = logging.getLogger(__name__)


class CUDService:
    """Kafka producer cud Service class for creating, updating, and deleting todo items."""

    TOPIC = "todos"

    @staticmethod
    def create_todo(data):
        """Create a new todo item."""
        result = todo_producer.send_message(data, CUDService.TOPIC, "create")
        if result:
            logger.info(f"Todo created: {data}")
        return result

    @staticmethod
    def update_todo(data):
        """Update an existing todo item."""
        result = todo_producer.send_message(data, CUDService.TOPIC, "update")
        if result:
            logger.info(f"Todo updated: {data}")
        return result

    @staticmethod
    def delete_todo(data):
        """Delete a todo item."""
        result = todo_producer.send_message(data, CUDService.TOPIC, "delete")
        if result:
            logger.info(f"Todo deleted: {data}")
        return result
