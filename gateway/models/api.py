from pydantic import BaseModel
from uuid import UUID

class API(BaseModel):
    _id: UUID
    name: str
    host: str
    key: str
    slug: str