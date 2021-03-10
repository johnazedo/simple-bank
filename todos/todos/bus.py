import lightbus

bus = lightbus.create()

def handle_new_user(event_message, username, email):
    print(f"A new user was created in the authentication service:")
    print(f"Username: {username}")
    print(f"Email: {email}")


print(bus.name)

@bus.client.on_start()
def bus_start():
    bus.auth.user.listen(
        handle_new_user,
        listener_name="print_on_new_registration"
    )

# valid = bus.auth.check_password(
#     username='admin',
#     password='secret'
# )

# if valid:
#     print('Authenticated')