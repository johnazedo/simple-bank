from fastapi.testclient import TestClient
from fastapi import status
from ..server import app

client = TestClient(app)

def test_create_api():
    response = client.post('/', json={
        'name':'Product',
        'host':'127.0.0.1',
        'key':'1512',
        'slug':'orders'
    })

    assert response.status_code == status.HTTP_201_CREATED