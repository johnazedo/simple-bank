from motor.motor_asyncio import AsyncIOMotorClient
from settings import Settings
USERPROFILE_DOC_TYPE = 'userprofile'


CLIENT = AsyncIOMotorClient(Settings().MONGO_URL)
DATABASE = CLIENT.gateway

def mongo_setup():
    collection = DATABASE.get_collection('api')
    collection.create_index('slug')