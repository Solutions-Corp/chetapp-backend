from pydantic_settings import BaseSettings
from functools import lru_cache
from dotenv import load_dotenv

load_dotenv(encoding="utf-8")


class Settings(BaseSettings):
    DATABASE_URL: str
    DEFAULT_USER_EMAIL: str
    DEFAULT_USER_PASSWORD: str
    JWT_SECRET_KEY: str
    JWT_ALGORITHM: str = "HS256"
    JWT_EXPIRES_MINUTES: int = 30
    APP_PORT: int = 8000

    class Config:
        env_file = ".env"


settings = Settings()


def get_settings() -> Settings:
    return settings
