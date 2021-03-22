from django.contrib import admin
from django.urls import path
from .views import ListTodos, DetailTodo, ListCategories, DetailCategory

urlpatterns = [
    path('', ListTodos.as_view(), name='list-create-todos'),
    path('<int:pk>', DetailTodo.as_view(), name='detail-todos'),
    path('categories/', ListCategories.as_view(), name='list-create-category'),
    path('category/<int:pk>', DetailCategory.as_view(), name='detail-category')
]
