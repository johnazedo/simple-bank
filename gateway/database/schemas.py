from uuid import UUID

from bson import ObjectId
from pydantic.main import BaseModel, Optional


class APISchema(BaseModel):
    _id: Optional[ObjectId]
    name: str
    host: str
    key: str
    slug: str

    class Config:
        schema_extra = {
            'example': {
                '_id': '1',
                'name': 'products',
                'host': '127.0.0.1',
                'key':'1',
                'slug':'products'
            }
        }