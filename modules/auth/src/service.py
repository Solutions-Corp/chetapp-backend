from src.utils import hash_password, verify_password, create_access_token, verify_access_token
from fastapi import HTTPException, Header
from src.models import User, UserRole
from sqlalchemy.orm import Session
from src.schemas import UserLogin
from typing import Optional
from src.config import settings
from datetime import datetime
import uuid


def create_default_user(db: Session, email: str):
    if not db.query(User).filter(User.email == email).first():
        default_user = User(
            id=uuid.uuid4(),
            email=settings.DEFAULT_USER_EMAIL,
            first_name="ChetApp",
            last_name="Admin",
            hashed_password=hash_password(settings.DEFAULT_USER_PASSWORD),
            role=UserRole.ADMIN,
            is_active=True,
            created_at=datetime.now(),
            last_login=None
        )
        db.add(default_user)
        db.commit()


def authenticate_user(user_data: UserLogin, db: Session) -> str:
    user = db.query(User).filter(User.email == user_data.email).first()
    if not user or not verify_password(user_data.password, user.hashed_password):
        raise HTTPException(status_code=401, detail="Invalid credentials")

    user.last_login = datetime.now()
    db.commit()
    return create_access_token({"sub": str(user.id), "email": user.email, "role": user.role})


def verify_token_from_header(authorization: Optional[str] = Header(None)):
    if not authorization:
        raise HTTPException(status_code=401, detail="Authorization header missing")
    
    scheme, _, token = authorization.partition(" ")
    if scheme.lower() != "bearer":
        raise HTTPException(status_code=401, detail="Invalid authentication scheme")
    if not token:
        raise HTTPException(status_code=401, detail="Token missing")
    
    payload = verify_access_token(token)
    if not payload:
        raise HTTPException(status_code=401, detail="Invalid token")
    return payload