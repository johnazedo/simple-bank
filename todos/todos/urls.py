from django.contrib import admin
from django.urls import path
from .views import ListTodo, UpdateTodo

urlpatterns = [
    path('', ListTodo.as_view(), name='list-create-todos'),
    path('<int:pk>', UpdateTodo.as_view(), name='update-todos')
]
