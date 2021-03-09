from django.contrib import admin
from django.urls import path
from .views import ListTodo, DetailTodo

urlpatterns = [
    path('', ListTodo.as_view(), name='list-todos'),
    path('<int:pk>', DetailTodo.as_view(), name='detail-todos')
]
