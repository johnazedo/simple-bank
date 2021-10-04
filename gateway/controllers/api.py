import uuid
from fastapi import APIRouter, status, Response, Depends
from gateway.models.api import API
from gateway.repositories.api import APIRepository

router = APIRouter()

async def routing(respository: APIRepository):
    pass

@router.get('/')
async def get():
    return {'Method': 'GET'}

@router.post('/')
async def post():
    return {'Method': 'POST'}

@router.put('/')
async def put():
    return {'Method': 'GET'}

@router.patch('/')
async def patch():
    return {'Method': 'POST'}

@router.delete('/')
async def delete():
    return {'Method': 'GET'}


@router.post('/create/api', status_code=status.HTTP_201_CREATED)
async def create_api(api: API, response: Response, repository = Depends(APIRepository)):
    result, err = await repository.create(api=api)
    if err:
        response.status_code = status.HTTP_403_FORBIDDEN
        return {'message':'API already exists.'}
    return result
