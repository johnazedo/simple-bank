from typing import List
from bson import ObjectId
from pymongo.results import InsertOneResult
from models.api import API
from database.db import DATABASE
from pymongo.errors import DuplicateKeyError
from typing import Tuple


class Repository:
    COLLECTION_NAME: str

    def __init__(self):
        self.get_collection()

    def get_collection(self):
        self.collection = DATABASE.get_collection(self.COLLECTION_NAME)


class APIRepository(Repository):
    COLLECTION_NAME = 'api'

    async def create(self, api: API) -> Tuple[API, bool]:
        try:
            result: InsertOneResult = await self.collection.insert_one(api.dict())
            api = await self.retrieve(_id=result.inserted_id)
            return (api, False)
        except DuplicateKeyError:
            return (None, True)
            
    async def retrieve(self, **kwargs) -> API:
        result = await self.collection.find_one(kwargs)
        return API(**result)

    async def fetch_all(self) -> List[API]:
        apis = [API(**result) async for result in self.collection.find()]
        return apis

    
    
