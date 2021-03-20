from django.shortcuts import render
from rest_framework.generics import RetrieveUpdateDestroyAPIView, ListCreateAPIView
from rest_framework.response import Response
from todos.models import Todo
from todos.serializers import TodoSerializer
from rest_framework.permissions import AllowAny


class ListTodo(ListCreateAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer


class UpdateTodo(RetrieveUpdateDestroyAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer
    lookup_field = 'pk'



