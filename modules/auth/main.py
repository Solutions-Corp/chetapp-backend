from src.routes.v1 import auth_v1_router
from src.service import create_default_user
from contextlib import asynccontextmanager
from src.database import Base, engine
from src.config import settings
from fastapi import FastAPI
from src.deps import get_db


app = FastAPI()


@app.on_event("startup")
async def startup():
    db = next(get_db()) 
    create_default_user(db, settings.DEFAULT_USER_EMAIL)


Base.metadata.create_all(bind=engine)

app.include_router(auth_v1_router, prefix="/api/v1/auth", tags=["auth"])
