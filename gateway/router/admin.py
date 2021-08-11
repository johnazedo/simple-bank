from django.contrib import admin
from .models import API

@admin.register(API)
class APIAdmin(admin.ModelAdmin):
    pass