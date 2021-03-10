from bus import bus

print("New user creation")
new_username = input("Enter a username: ").strip()
new_email = input("Enter the user's email address: ").strip()

# You would normally store the new user in your database
# at this point. We don't show this here for simplicity.

# Let the bus know a user has been registered by firing the event
bus.auth.user.fire(
    username=new_username,
    email=new_email
)

print("Done")