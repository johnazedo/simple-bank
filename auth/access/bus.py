# Testing lightbus

import lightbus

bus = lightbus.create()


class AuthApi(lightbus.Api):
    user = lightbus.Event(parameters=('username', 'email'))

    class Meta:
        name = 'auth'
    
    def check_password(self, username, password):
        return username == 'admin' and password == 'secret'

bus.client.register_api(AuthApi())