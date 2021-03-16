from django.shortcuts import render
from rest_framework.generics import ListAPIView, RetrieveAPIView
from rest_framework.response import Response
from todos.models import Todo
from todos.serializers import TodoSerializer
from rest_framework.permissions import AllowAny


class ListTodo(ListAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer


class DetailTodo(RetrieveAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer
    lookup_field = 'pk'

