from django.urls import path
from .views import Router
namespace='router'

urlpatterns = [
    path('', Router.as_view(), name='all')
]
