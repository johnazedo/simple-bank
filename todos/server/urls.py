from django.contrib import admin
from django.urls import path, include
from todos.urls import categories

urlpatterns = [
    path('admin/', admin.site.urls),
    path('todos/', include('todos.urls')),
    path('categories/', include(categories))
]
