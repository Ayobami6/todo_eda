from django.urls import path
from .views import TodoCreateAPIView

urlpatterns = [
    path("", TodoCreateAPIView.as_view(), name="create-todo"),
]
