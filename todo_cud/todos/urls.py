from django.urls import path
from .views import TodoAPIView

urlpatterns = [
    path("", TodoAPIView.as_view(), name="create-todo"),
    path("", TodoAPIView.as_view(), name="list-todo"),
]
