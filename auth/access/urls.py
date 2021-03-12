from django.contrib import admin
from django.urls import path
from .views import SessionTest, SignUpView

urlpatterns = [
    path('<int:pk>', SessionTest.as_view(), name='list-users'),
    path('signup/', SignUpView.as_view(), name='register')
]
