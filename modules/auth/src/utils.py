from passlib.context import CryptContext
from datetime import datetime, timedelta
from src.config import settings
from jose import jwt
import importlib
import bcrypt
import sys

if not hasattr(bcrypt, "__about__"):

    class _About:
        __version__ = getattr(bcrypt, "__version__", "unknown")

    bcrypt.__about__ = _About()


if "passlib.handlers.bcrypt" in sys.modules:
    importlib.reload(sys.modules["passlib.handlers.bcrypt"])


pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")


def hash_password(password: str) -> str:
    return pwd_context.hash(password)


def verify_password(plain: str, hashed: str) -> bool:
    return pwd_context.verify(plain, hashed)


def create_access_token(data: dict) -> str:
    expire = datetime.now() + timedelta(minutes=settings.JWT_EXPIRES_MINUTES)
    data.update({"exp": expire})
    return jwt.encode(data, settings.JWT_SECRET_KEY, algorithm=settings.JWT_ALGORITHM)
