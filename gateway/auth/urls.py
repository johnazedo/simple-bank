from django.contrib import admin
from django.urls import path
from .views import SessionTest, SignUpView

urlpatterns = [
    path('signup/', SignUpView.as_view(), name='register')
]
