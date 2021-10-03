from fastapi.testclient import TestClient
from fastapi import status
from gateway.server import app
from repositories.api import APIRepository

class APIRepositoryTestCase(APIRepository):
    COLLECTION_NAME = 'test_api'

    def setup(self):
        self.collection.create_index('slug', unique=True)

    def drop_collection(self):
        self.collection.drop()


app.dependency_overrides[APIRepository] = APIRepositoryTestCase
client = TestClient(app)

class TestAPI:
    def setup_class(self):
        self.repository = APIRepositoryTestCase()
        self.repository.setup()

    def teardown_class(self):
        self.repository.drop_collection()

    def test_create_api(self):
        response = client.post('/', json={
            'name':'Product',
            'host':'127.0.0.1',
            'key':'1512',
            'slug':'orders'
        })
        assert response.status_code == status.HTTP_201_CREATED
