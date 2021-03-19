from django.db import models
from django.contrib.auth.models import User

# Create your models here.
class Todo(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE, null=True)
    text = models.CharField(max_length=512, null=False)
    done = models.BooleanField(default=False)

    class Meta:
        verbose_name = 'Todo'
        verbose_name_plural = 'Todos'
    
    def __str__(self):
        return self.text