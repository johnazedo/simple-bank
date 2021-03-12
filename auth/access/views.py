from django.shortcuts import render
from rest_framework.generics import ListAPIView, RetrieveAPIView, CreateAPIView
from rest_framework.response import Response
from rest_framework.permissions import AllowAny, IsAuthenticated
from .serializers import UserSerializer, SignUpSerializer
from django.contrib.auth.models import User

# Create your views here.

class SessionTest(RetrieveAPIView):
    permission_classes = [IsAuthenticated]
    serializer_class = UserSerializer
    queryset = User.objects.all()
    lookup_field = 'pk'


class SignUpView(CreateAPIView):
    serializer_class = SignUpSerializer
