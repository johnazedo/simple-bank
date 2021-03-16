from django.db import models
from django.dispatch import receiver
from django.db.models.signals import post_save
from django.contrib.auth.models import User
import redis as rd
# Create your models here.

@receiver(post_save, sender=User, dispatch_uid="on_user_save")
def on_user_save(sender, instance: User, created, **kwargs):
    redis = rd.StrictRedis(host='127.0.0.1', port=6379, db=0)
    redis.mset({
        'username': instance.username,
        'password': instance.password
    })

    if created:
        pass