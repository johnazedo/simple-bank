from django.urls import path, include
from .views import SignUpView, UserView


users = (
    [
        path('/', UserView.as_view(), name='list')

    ], 'users'
)

urlpatterns = [
    path('signup', SignUpView.as_view(), name='register'),
    path('users', include(users))
]
