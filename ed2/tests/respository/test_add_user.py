from repositories import UserRepository
from user import User

def test_add_user(empty_repository: UserRepository) -> None:
    user = User("John", "Doe", 1990, "user", 0)
    assert empty_repository.add_user(user) == True
