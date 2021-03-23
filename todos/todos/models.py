from django.db import models
from django.contrib.auth.models import User

# Create your models here.
class Todo(models.Model):
    category = models.ForeignKey('Category', on_delete=models.CASCADE, null=True, blank=True)
    user = models.ForeignKey(User, on_delete=models.CASCADE, null=True)
    text = models.CharField(max_length=512, null=False)
    done = models.BooleanField(default=False)

    class Meta:
        verbose_name = 'Todo'
        verbose_name_plural = 'Todos'
    
    def __str__(self):
        return self.text


class Category(models.Model):
    name = models.CharField(max_length=128, null=False)
    color = models.CharField(max_length=6, null=False)
    icon_code = models.PositiveIntegerField(null=False, blank=False)

    class Meta:
        verbose_name = 'Category'
        verbose_name_plural = 'Categories'

    def __str__(self):
        return self.name