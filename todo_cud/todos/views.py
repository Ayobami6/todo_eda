from django.shortcuts import render
from rest_framework.views import APIView
from sparky_utils.response import service_response
from .serializers import TodoCreateSerializer, TodoSerializer
from .kafka_producer import todo_producer
from sparky_utils.advice import exception_advice
from .models import Todo

# Create your views here.


class RootAPIView(APIView):
    def get(self, request):
        return service_response(message="Todo CUD Service is up and running.")


class TodoAPIView(APIView):

    @exception_advice()
    def post(self, request):
        serializer = TodoCreateSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        data = serializer.validated_data
        producer_topic = "todos"
        # send event
        todo_producer.send_message(data, producer_topic, "create")
        return service_response(
            status="success", message="Todo create successfully queued", status_code=202
        )

    @exception_advice()
    def get(self, request, *args, **kwargs):
        todos = Todo.objects.all()
        # serialize the todos
        serializer = TodoSerializer(todos, many=True)

        return service_response(
            message="Todos fetch successfully",
            data=serializer.data,
            status_code=200,
        )

    @exception_advice()
    def delete(self, request, *args, **kwargs):
        todo_id = kwargs.get("id")
        producer_topic = "todos"
        # send event
        todo_producer.send_message({"id": todo_id}, producer_topic, "delete")
        return service_response(
            status="success", message="Todo delete successfully queued", status_code=202
        )

    @exception_advice()
    def patch(self, request, *args, **kwargs):
        todo_id = kwargs.get("id")
        serializer = TodoCreateSerializer(data=request.data, partial=True)
        serializer.is_valid(raise_exception=True)
        data = serializer.validated_data
        data["id"] = todo_id
        producer_topic = "todos"
        # send event
        todo_producer.send_message(data, producer_topic, "update")
        return service_response(
            status="success", message="Todo update successfully queued", status_code=202
        )
