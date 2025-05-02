from sqlalchemy import Column, Boolean, String, UUID, DateTime, Enum
from src.database import Base
import enum


class UserRole(str, enum.Enum):
    ADMIN = "ADMIN"
    OPERATOR = "OPERATOR"


class User(Base):
    __tablename__ = "users"

    id = Column(UUID(as_uuid=True), primary_key=True, index=True)
    first_name = Column(String, nullable=False)
    last_name = Column(String, nullable=False)
    email = Column(String, unique=True, index=True)
    hashed_password = Column(String, nullable=False)
    is_active = Column(Boolean, default=True)
    created_at = Column(DateTime, nullable=False)
    last_login = Column(DateTime, nullable=True)
    role = Column(Enum(UserRole), default=UserRole.OPERATOR, nullable=False)
