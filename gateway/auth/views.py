from django.shortcuts import render
from rest_framework.generics import ListAPIView, RetrieveAPIView, CreateAPIView
from rest_framework.response import Response
from rest_framework.permissions import AllowAny, IsAuthenticated
from .serializers import UserSerializer, SignUpSerializer
from django.contrib.auth.models import User

# Create your views here.
class SignUpView(CreateAPIView):
    serializer_class = SignUpSerializer


