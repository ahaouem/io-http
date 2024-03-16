from repositories import UserRepository
from user import User

def test_patch_user_failure(filled_repository: UserRepository) -> None:
    new_user = User("John", "Doe", 1990, "user", 0)
    assert filled_repository.patch_user(5, new_user) == False
    assert filled_repository.get_user(0) != new_user
