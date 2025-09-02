from django.db import models

# Create your models here.


class Todo(models.Model):
    title = models.CharField(max_length=100)
    description = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    completed = models.BooleanField(default=False)

    def __str__(self):
        return self.title

    class Meta:
        ordering = ["-created_at"]
        verbose_name = "Todo"
        verbose_name_plural = "Todos"
        db_table = "todos"
        indexes = [models.Index(fields=["title"])]
