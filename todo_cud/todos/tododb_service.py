import logging
from .models import Todo


class TodoDBService:
    """Service class for handling database operations for todo items."""

    @staticmethod
    def process_message(message):
        """Process a message to perform the corresponding database operation."""
        operation = message.get("operation")
        data = message.get("data")

        if operation == "create":
            return TodoDBService.create_todo_db(data)
        elif operation == "update":
            return TodoDBService.update_todo_db(data)
        elif operation == "delete":
            return TodoDBService.delete_todo_db(data)
        else:
            logging.error(f"Unknown operation: {operation}")
            return None

    @staticmethod
    def create_todo_db(data):
        """Create a new todo item in the database."""

        try:
            todo = Todo.objects.create(**data)
            logging.info(f"Todo created in DB: {todo}")
            return todo
        except Exception as e:
            logging.error(f"Failed to create todo in DB: {e}")
            return None

    @staticmethod
    def update_todo_db(data):
        """Update an existing todo item in the database."""
        try:
            todo = Todo.objects.get(id=data["id"])
            for key, value in data.items():
                setattr(todo, key, value)
            todo.save()
            logging.info(f"Todo updated in DB: {todo}")
            return todo
        except Todo.DoesNotExist:
            logging.error(f"Todo with id {data['id']} does not exist.")
            return None
        except Exception as e:
            logging.error(f"Failed to update todo in DB: {e}")
            return None

    @staticmethod
    def delete_todo_db(data):
        """Delete a todo item from the database."""
        try:
            todo = Todo.objects.get(id=data["id"])
            todo.delete()
            logging.info(f"Todo deleted from DB: {todo}")
            return todo
        except Todo.DoesNotExist:
            logging.error(f"Todo with id {data['id']} does not exist.")
            return None
        except Exception as e:
            logging.error(f"Failed to delete todo from DB: {e}")
            return None
