from repositories import UserRepository

def test_get_user_failure(filled_repository: UserRepository, empty_repository: UserRepository) -> None:
    assert empty_repository.get_user(0) == None
    assert filled_repository.get_user(5) == None
