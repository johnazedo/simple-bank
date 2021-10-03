from typing import List
from bson import ObjectId
from pymongo.results import InsertOneResult
from models.api import API
from database.db import DATABASE
from pymongo.errors import DuplicateKeyError
from typing import Tuple


class APIRepository:
    COLLECTION_NAME = 'api'

    def __init__(self):
        self.collection = DATABASE.get_collection(self.COLLECTION_NAME)

    async def create(self, api: API) -> Tuple[API, bool]:
        try:
            result: InsertOneResult = await self.collection.insert_one(api.dict())
            return (await self.retrieve(_id=result.inserted_id), False)
        except DuplicateKeyError:
            return (None, True)
            
    async def retrieve(self, **kwargs) -> API:
        result = await self.collection.find_one(kwargs)
        return API(**result)

    async def fetch_all(self) -> List[API]:
        apis = [API(**result) async for result in self.collection.find()]
        return apis