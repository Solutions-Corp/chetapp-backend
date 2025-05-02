from src.service import authenticate_user, verify_token_from_header
from fastapi import APIRouter, Depends
from src.schemas import UserLogin, Token, TokenPayload
from sqlalchemy.orm import Session
from src.deps import get_db

auth_v1_router = APIRouter()

@auth_v1_router.post("/login", response_model=Token)
def login(user: UserLogin, db: Session = Depends(get_db)):  
    token = authenticate_user(user, db)
    return {"access_token": token, "token_type": "bearer"}


@auth_v1_router.get("/verify", response_model=TokenPayload)
def verify(payload: dict = Depends(verify_token_from_header)):
    return payload