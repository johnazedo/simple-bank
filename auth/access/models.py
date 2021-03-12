from django.db import models
from django.dispatch import receiver
from django.db.models.signals import post_save
from django.contrib.auth.models import User
# Create your models here.

@receiver(post_save, sender=User, dispatch_uid="on_user_save")
def on_user_save(sender, instance, created, **kwargs):
    if created:
        print('Testando')