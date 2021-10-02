from fastapi import FastAPI
from controllers.api import router as api_router
from database.db import mongo_setup

app = FastAPI()

@app.on_event('startup')
async def startup():
    mongo_setup()

app.include_router(api_router)