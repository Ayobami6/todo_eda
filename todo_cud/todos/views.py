from django.shortcuts import render
from rest_framework.views import APIView
from sparky_utils.response import service_response

# Create your views here.


class RootAPIView(APIView):
    def get(self, request):
        return service_response(message="Todo CUD Service is up and running.")


class TodoCreateAPIView(APIView):
    def post(self, request):
        pass
