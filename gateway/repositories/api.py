from typing import List
from bson import ObjectId
from pymongo.results import InsertOneResult
from models.api import API
from database.db import DATABASE


class APIRepository:
    COLLECTION_NAME = 'api'

    def __init__(self):
        self.collection = DATABASE.get_collection(self.COLLECTION_NAME)

    async def add(self, api: API) -> API:
        result: InsertOneResult = await self.collection.insert_one(api.dict())
        return await self.retrieve_by_id(id=result.inserted_id)

    async def retrieve_by_slug(self, slug: str) -> API:
        pass

    async def retrieve_by_id(self, id: ObjectId) -> API:
        result = await self.collection.find_one({'_id': id})
        return API(**result)

    async def retrieve(self) -> List[API]:
        apis: List[API] = []
        async for result in self.collection.find():
            apis.append(API(**result))
        return apis