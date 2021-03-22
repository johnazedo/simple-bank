from django.contrib import admin
from .models import Todo, Category

@admin.register(Todo)
class TodoAdmin(admin.ModelAdmin):
    pass


@admin.register(Category)
class CategoryAdmin(admin.ModelAdmin):
    pass