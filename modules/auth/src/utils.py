from passlib.context import CryptContext
from datetime import datetime, timedelta
from jose import jwt, JWTError


def load_private_key_pem() -> bytes:
    return open("keys/private.pem", "rb").read()

def load_public_key_pem() -> bytes:
    return open("keys/public.pem", "rb").read()


pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

def hash_password(password: str) -> str:
    return pwd_context.hash(password)

def verify_password(plain: str, hashed: str) -> bool:
    return pwd_context.verify(plain, hashed)

def create_access_token(data: dict) -> str:
    expire = datetime.now() + timedelta(minutes=15)
    data.update({"exp": expire})
    private_pem = load_private_key_pem() 
    token = jwt.encode(data, private_pem, algorithm="RS256")
    return token

def verify_access_token(token: str) -> dict | None:
    public_pem = load_public_key_pem()
    try:
        payload = jwt.decode(token, public_pem, algorithms=["RS256"])
        return payload
    except JWTError:
        return None

def verify_access_token(token: str) -> dict | None:
    public_pem = load_public_key_pem()
    try:
        payload = jwt.decode(token, public_pem, algorithms=["RS256"])
        return payload
    except JWTError:
        return None