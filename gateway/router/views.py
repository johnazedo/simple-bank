from typing import Dict
from rest_framework.authentication import TokenAuthentication
from rest_framework.views import APIView
from django.http import JsonResponse

from .models import API
from rest_framework import status
from rest_framework.permissions import IsAuthenticated, AllowAny


# Create your views here.
class Router(APIView):
    permission_class = [AllowAny, ]
    authentication_classes = [TokenAuthentication, ]
    model = API

    def dispatch(self, request, *args, **kwargs):
        name = self.get_api_name(self.request.path)
        api = self.model.objects.filter(name=name).first()

        if api is None:
            return InvalidRequest(msg='Not found!', status_code=status.HTTP_404_NOT_FOUND)

        data = api.send_request(
            path=self.request.path,
            method=self.request.method
        )

        return JsonResponse(data)


    def get_api_name(self, path: str) -> str:
        path_array = path.split('/')
        return path_array[0]


class InvalidRequest:
    msg: str
    status_code: int

    def __init__(self, msg, status_code):
        self.msg = msg
        self.status_code = status_code

    def to_dict(self) -> Dict:
        return {
            'msg': self.msg,
            'status_code': self.status_code
        }

    def __str__(self) -> JsonResponse:
        return JsonResponse(self.to_dict())