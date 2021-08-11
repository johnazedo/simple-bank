from django.db import models

# Create your models here.
class API(models.Model):
    name = models.CharField(max_length=255)
    localhost = models.CharField(max_length=255)
    key = models.CharField(max_length=255)
    slug = models.SlugField(max_length=255)

    class Meta:
        verbose_name='Server'
        verbose_name_plural='Servers'

    def __str__(self):
        return self.name