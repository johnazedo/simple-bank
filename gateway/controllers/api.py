import uuid
from fastapi import APIRouter, status
from models.api import API
from repositories.api import APIRepository

router = APIRouter()

@router.get('/', status_code=status.HTTP_201_CREATED)
async def gate():
    api = API(
        name='Orders',
        host='127.0.0.1',
        key='1512',
        slug='orders'
    )
    repository = APIRepository()
    result = await repository.add(api=api)
    # result = await repository.retrieve()
    return result