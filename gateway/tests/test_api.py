import pytest
from fastapi.testclient import TestClient
from fastapi import status
from gateway.server import app
from gateway.repositories.api import APIRepository


class APIRepositoryTestCase(APIRepository):
    COLLECTION_NAME = 'test_api'

    def drop_collection(self):
        self.collection.drop()


@pytest.fixture
def client():
    app.dependency_overrides[APIRepository] = APIRepositoryTestCase
    client = TestClient(app)
    yield client
    APIRepositoryTestCase().drop_collection()


def test_create_api(client):
    response = client.post('/create/api', json={
        'name':'Product',
        'host':'127.0.0.1',
        'key':'1512',
        'slug':'orders'
    })
    assert response.status_code == status.HTTP_201_CREATED



    
