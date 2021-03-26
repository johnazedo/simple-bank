from django.shortcuts import render
from rest_framework.generics import RetrieveUpdateDestroyAPIView, ListCreateAPIView
from rest_framework.response import Response
from todos.models import Todo, Category
from todos.serializers import TodoSerializer, CategorySerializer
from rest_framework.permissions import AllowAny
from django_filters.rest_framework.backends import DjangoFilterBackend
from todos.filters import TodoFilter

class ListTodos(ListCreateAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    filter_backends = (DjangoFilterBackend,)
    filterset_class = TodoFilter
    serializer_class = TodoSerializer


class DetailTodo(RetrieveUpdateDestroyAPIView):
    permission_classes = [AllowAny]
    queryset = Todo.objects.all()
    serializer_class = TodoSerializer
    lookup_field = 'pk'


class ListCategories(ListCreateAPIView):
    permission_classes = [AllowAny]
    queryset = Category.objects.all()
    serializer_class = CategorySerializer


class DetailCategory(RetrieveUpdateDestroyAPIView):
    permission_classes = [AllowAny]
    queryset = Category.objects.all()
    serializer_class = CategorySerializer
    lookup_field = 'pk'



