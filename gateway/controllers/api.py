import uuid
from fastapi import APIRouter, status, Response, Depends
from models.api import API
from repositories.api import APIRepository

router = APIRouter()

@router.post('/', status_code=status.HTTP_201_CREATED)
async def create_api(api: API, response: Response, repository = Depends(APIRepository)):
    result, err = await repository.create(api=api)
    if err:
        response.status_code = status.HTTP_403_FORBIDDEN
        return {'message':'API already exists.'}
    return result
