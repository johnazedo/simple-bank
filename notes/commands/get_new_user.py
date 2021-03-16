from django.core.management.base import BaseCommand, CommandError
import redis as rd
from django.contrib.auth.models import User
class Command(BaseCommand):
    def handle(self, *args, **options):
        redis = rd.StrictRedis(host='127.0.0.1', port=6379, db=0)
        blocked_list = []

        while True:
            key = 'username'
            if redis.exists(key):
                username = redis.get('username').decode('utf-8')
                password = redis.get('password').decode('utf-8')
                user = User.objects.create_user(username, password)
                blocked_list.append(key)
                
            if key in blocked_list:
                break