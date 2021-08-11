from typing import Dict
from django.db import models
import requests

# Create your models here.
from django.http import HttpResponse
from requests import Response


class API(models.Model):
    name = models.CharField(max_length=255)
    localhost = models.CharField(max_length=255)
    key = models.CharField(max_length=255)
    slug = models.SlugField(max_length=255)

    class Meta:
        verbose_name='API'
        verbose_name_plural='APIs'

    def __str__(self):
        return self.name


    def get_absolute_url(self, path: str) -> str:
        return f'http://{self.localhost}/'

    def send_request(self, path: str, method: str) -> HttpResponse:
        url = self.get_absolute_url(path=path)
        result = requests.Response()

        if method == 'GET':
            result = requests.get(url)

        return HttpResponse(
            content=result.content,
            status=result.status_code,
            content_type=result.headers,
        )