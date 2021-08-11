from rest_framework.generics import ListAPIView, CreateAPIView
from rest_framework.permissions import IsAdminUser
from .serializers import UserSerializer, SignUpSerializer
from rest_framework.authentication import TokenAuthentication

# Create your views here.
class SignUpView(CreateAPIView):
    serializer_class = SignUpSerializer


class UserView(ListAPIView):
    serializer_class = UserSerializer
    permission_classes = [IsAdminUser,]
    authentication_classes = [TokenAuthentication,]
