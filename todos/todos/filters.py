from django_filters.rest_framework import FilterSet, filters
from todos.models import Todo

class TodoFilter(FilterSet):
    class Meta:
        model = Todo
        fields = ['category']