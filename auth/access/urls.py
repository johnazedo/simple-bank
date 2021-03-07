from django.contrib import admin
from django.urls import path
from .views import SessionTest

urlpatterns = [
    path('<int:pk>', SessionTest.as_view(), name='list-users'),
]
