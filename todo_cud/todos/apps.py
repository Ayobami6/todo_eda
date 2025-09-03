from django.apps import AppConfig
from .consumer import todo_consumer


class TodosConfig(AppConfig):
    default_auto_field = "django.db.models.BigAutoField"
    name = "todos"

    # def ready(self):
    #     todo_consumer.start()
